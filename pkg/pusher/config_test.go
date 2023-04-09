// Copyright 2023 Dimitri Koshkin. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package pusher

import (
	"bytes"
	"testing"

	"github.com/sethvargo/go-githubactions"
	"github.com/stretchr/testify/assert"
)

func TestNewFromInputs(t *testing.T) {
	// ...
	actionLog := bytes.NewBuffer(nil)
	keys := "workflow=checks,job=status-writer-action"
	envMap := map[string]string{
		"INPUT_KEYS": keys,
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
		"job":      "status-writer-action",
	}
	assert.Equal(t, cfg.Keys, expected)
}
