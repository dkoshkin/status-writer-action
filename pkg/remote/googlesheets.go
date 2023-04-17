// Copyright 2023 Dimitri Koshkin. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package remote

import (
	"context"
	"fmt"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

//nolint:gochecknoglobals // a list of headers for the spreadsheet
var headers = []string{"Timestamp", "Actor", "Status"}

type GoogleSheetsConfig struct {
	SpreadsheetID string
}

type GoogleSheetsWriter struct {
	service     *sheets.Service
	spreadsheet *sheets.Spreadsheet
}

func NewGoogleSheetsWriter(ctx context.Context, cfg *GoogleSheetsConfig) (Writer, error) {
	// create new client using default credentials
	client, err := google.DefaultClient(ctx, sheets.SpreadsheetsScope)
	if err != nil {
		return nil, fmt.Errorf("error creating Google Sheets client: %w", err)
	}

	// create new service using client
	service, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("error creating Google Sheets service: %w", err)
	}

	spreadsheet, err := service.Spreadsheets.Get(cfg.SpreadsheetID).Do()
	if err != nil {
		return nil, fmt.Errorf("error getting spreadsheet by ID: %w", err)
	}

	return &GoogleSheetsWriter{service: service, spreadsheet: spreadsheet}, nil
}

func (w *GoogleSheetsWriter) Write(ctx context.Context, data Data) error {
	sheet, err := w.ensureSheetExists(ctx, data)
	if err != nil {
		return fmt.Errorf("error ensuring sheet exists: %w", err)
	}

	row := &sheets.ValueRange{
		Values: [][]interface{}{{}},
	}
	for _, val := range []string{timestampAsString(), data.Actor, data.Status} {
		row.Values[0] = append(row.Values[0], val)
	}
	for _, tag := range data.Tags {
		row.Values[0] = append(row.Values[0], tag.Value)
	}

	err = appendRow(ctx, w.service, w.spreadsheet.SpreadsheetId, sheet.Properties.Title, row)
	if err != nil {
		return fmt.Errorf("error addings data to sheet: %w", err)
	}

	return nil
}

func (w *GoogleSheetsWriter) ensureSheetExists(
	ctx context.Context,
	data Data,
) (*sheets.Sheet, error) {
	var sheet *sheets.Sheet
	for n := range w.spreadsheet.Sheets {
		sheetPtr := w.spreadsheet.Sheets[n]
		if sheetPtr != nil && sheetPtr.Properties.Title == data.Repository {
			sheet = sheetPtr
		}
	}

	// sheet already exists, return it
	if sheet != nil {
		return sheet, nil
	}

	// create a new sheet
	sheet = &sheets.Sheet{
		Properties: &sheets.SheetProperties{
			Title: data.Repository,
		},
	}
	w.spreadsheet.Sheets = append(w.spreadsheet.Sheets, sheet)

	// create a new sheet in the spreadsheet
	err := createSheet(ctx, w.service, w.spreadsheet.SpreadsheetId, sheet.Properties.Title)
	if err != nil {
		return nil, fmt.Errorf("error creating new sheet: %w", err)
	}

	row := &sheets.ValueRange{
		Values: [][]interface{}{{}},
	}
	// add headers on the first row
	for _, header := range headers {
		row.Values[0] = append(row.Values[0], header)
	}
	// add additional headers from tags
	// this assumes that the tags are the same across all write requests
	// TODO support dynamic tags
	for _, tag := range data.Tags {
		row.Values[0] = append(row.Values[0], tag.Key)
	}

	err = appendRow(ctx, w.service, w.spreadsheet.SpreadsheetId, sheet.Properties.Title, row)
	if err != nil {
		return nil, fmt.Errorf("error adding headers to sheet: %w", err)
	}

	return sheet, nil
}

func createSheet(
	ctx context.Context,
	service *sheets.Service,
	spreadsheetID string,
	sheetName string,
) error {
	_, err := service.Spreadsheets.BatchUpdate(
		spreadsheetID,
		&sheets.BatchUpdateSpreadsheetRequest{
			Requests: []*sheets.Request{
				{
					AddSheet: &sheets.AddSheetRequest{
						Properties: &sheets.SheetProperties{
							Title: sheetName,
						},
					},
				},
			},
		},
	).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("error creating new sheet: %w", err)
	}

	return nil
}

func appendRow(
	ctx context.Context,
	service *sheets.Service,
	spreadsheetID string,
	sheetName string,
	row *sheets.ValueRange,
) error {
	_, err := service.Spreadsheets.Values.
		Append(spreadsheetID, sheetName, row).
		ValueInputOption("USER_ENTERED").
		InsertDataOption("INSERT_ROWS").
		Context(ctx).
		Do()
	if err != nil {
		return fmt.Errorf("error addings data to sheet: %w", err)
	}

	return nil
}
