// Copyright 2023 Dimitri Koshkin. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package remote

import (
	"context"
	"fmt"

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

	org    string
	bucket string
}

func NewInfluxDBWriter(cfg *InfluxDBConfig) Writer {
	return &InfluxDBWriter{
		client: influxdb2.NewClient(cfg.URL, cfg.Token),
		org:    cfg.Org,
		bucket: cfg.Bucket,
	}
}

func (w InfluxDBWriter) Write(ctx context.Context, data Data) error {
	writeAPI := w.client.WriteAPIBlocking(w.org, w.bucket)

	fields := map[string]interface{}{
		"actor":  data.Actor,
		"status": data.Status,
	}

	tagsMap := make(map[string]string)
	for _, tag := range data.Tags {
		tagsMap[tag.Key] = tag.Value
	}
	point := influxdb2.NewPoint(data.Repository, tagsMap, fields, timestamp())

	if err := writeAPI.WritePoint(ctx, point); err != nil {
		return fmt.Errorf("error writing data to InfluxDB: %w", err)
	}

	return nil
}
