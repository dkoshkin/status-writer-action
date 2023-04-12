// Copyright 2023 Dimitri Koshkin. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package remote

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/Iwark/spreadsheet.v2"
)

//nolint:gochecknoglobals // a list of headers for the spreadsheet
var headers = []string{"Timestamp", "Actor", "Status"}

type GoogleSheetsConfig struct {
	SpreadsheetID string
}

type GoogleSheetsWriter struct {
	service     *spreadsheet.Service
	spreadsheet *spreadsheet.Spreadsheet
}

func NewGoogleSheetsWriter(cfg *GoogleSheetsConfig) (Writer, error) {
	service, err := spreadsheet.NewService()
	if err != nil {
		return nil, fmt.Errorf("error creating Google Sheets service: %w", err)
	}
	ss, err := service.FetchSpreadsheet(cfg.SpreadsheetID)
	if err != nil {
		return nil, fmt.Errorf("error fetching spreadsheet by ID: %w", err)
	}

	return &GoogleSheetsWriter{service: service, spreadsheet: &ss}, nil
}

func (w *GoogleSheetsWriter) Write(ctx context.Context, data Data) error {
	sheet, err := w.ensureSheetExists(data)
	if err != nil {
		return fmt.Errorf("error ensuring sheet exists: %w", err)
	}

	nextRow := len(sheet.Rows)
	for n, val := range []string{time.Now().String(), data.Actor, data.Status} {
		sheet.Update(nextRow, n, val)
	}
	for n, tag := range data.Tags {
		sheet.Update(nextRow, len(headers)+n, tag.Value)
	}

	err = sheet.Synchronize()
	if err != nil {
		return fmt.Errorf("error synchronizing sheet: %w", err)
	}

	return nil
}

func (w *GoogleSheetsWriter) ensureSheetExists(data Data) (*spreadsheet.Sheet, error) {
	sheet, err := w.spreadsheet.SheetByTitle(data.Repository)
	//nolint:nestif // don't know how to reduce the complexity here
	if err != nil {
		// return if the error is not for a sheet not found
		if err.Error() != "sheet not found by the title" {
			return nil, fmt.Errorf("error getting sheet by title: %w", err)
		}
		// create a new sheet if it doesn't exist
		err = w.service.AddSheet(w.spreadsheet, spreadsheet.SheetProperties{Title: data.Repository})
		if err != nil {
			return nil, fmt.Errorf("error creating a new sheet: %w", err)
		}
		sheet, err = w.spreadsheet.SheetByTitle(data.Repository)
		if err != nil {
			return nil, fmt.Errorf("error getting sheet by title: %w", err)
		}

		// add headers on the first row
		for n, header := range headers {
			sheet.Update(0, n, header)
		}
		// add additional headers from tags
		// this assumes that the tags are the same across all write requests
		// TODO support dynamic tags
		for n, tag := range data.Tags {
			sheet.Update(0, len(headers)+n, tag.Key)
		}

		err = sheet.Synchronize()
		if err != nil {
			return nil, fmt.Errorf("error synchronizing sheet: %w", err)
		}

		// get the sheet again after adding headers
		sheet, err = w.spreadsheet.SheetByTitle(data.Repository)
		if err != nil {
			return nil, fmt.Errorf("error getting sheet by title after adding headers: %w", err)
		}
	}

	return sheet, nil
}
