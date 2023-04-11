// Copyright 2023 Dimitri Koshkin. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package remote

import (
	"bytes"
	"testing"

	"github.com/sethvargo/go-githubactions"
	"github.com/stretchr/testify/assert"
)

func TestNewFromInputs(t *testing.T) {
	// ...
	actionLog := bytes.NewBuffer(nil)
	tags := "workflow=checks,job=build-and-run"
	envMap := map[string]string{
		"INPUT_BACKEND":         "influxdb",
		"INPUT_INFLUXDB_TOKEN":  "token",
		"INPUT_INFLUXDB_URL":    "http://localhost",
		"INPUT_INFLUXDB_ORG":    "org",
		"INPUT_INFLUXDB_BUCKET": "bucket",
		"INPUT_REPOSITORY":      "repo",
		"INPUT_ACTOR":           "actor",
		"INPUT_STATUS":          "success",
		"INPUT_TAGS":            tags,
	}
	getenv := func(key string) string {
		return envMap[key]
	}
	action := githubactions.New(
		githubactions.WithWriter(actionLog),
		githubactions.WithGetenv(getenv),
	)
	cfg, err := NewFromInputs(action)
	assert.NoError(t, err)
	assert.Equal(t, "", actionLog.String())
	expected := map[string]string{
		"workflow": "checks",
		"job":      "build-and-run",
	}
	assert.Equal(t, cfg.Data.Repository, "repo")
	assert.Equal(t, cfg.Data.Actor, "actor")
	assert.Equal(t, cfg.Data.Status, "success")
	assert.Equal(t, cfg.Data.Tags, expected)
}
