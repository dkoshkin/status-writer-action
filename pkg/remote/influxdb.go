// Copyright 2023 Dimitri Koshkin. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package remote

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxDBConfig struct {
	Token  string
	URL    string
	Org    string
	Bucket string
}

type InfluxDBWriter struct {
	client influxdb2.Client

	Org    string
	Bucket string
}

func NewInfluxDBWriter(cfg *InfluxDBConfig) Writer {
	return &InfluxDBWriter{
		client: influxdb2.NewClient(cfg.URL, cfg.Token),
		Org:    cfg.Org,
		Bucket: cfg.Bucket,
	}
}

func (w InfluxDBWriter) Write(ctx context.Context, data Data) error {
	writeAPI := w.client.WriteAPIBlocking(w.Org, w.Bucket)

	fields := map[string]interface{}{
		"actor":  data.Actor,
		"status": data.Status,
	}
	point := influxdb2.NewPoint(data.Repository, data.Tags, fields, time.Now())

	if err := writeAPI.WritePoint(ctx, point); err != nil {
		return fmt.Errorf("error writing data to InfluxDB: %w", err)
	}

	return nil
}
