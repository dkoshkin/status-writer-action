// Copyright 2023 Dimitri Koshkin. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package remote

import (
	"fmt"
	"strings"

	"github.com/sethvargo/go-githubactions"
)

// Backend enum type.
type Backend int

func (s Backend) String() string {
	//nolint:gocritic // Prefer switch statement over if statement.
	switch s {
	case BackendInfluxDB:
		return "influxdb"
	}
	return "unknown"
}

func StringToBackend(in string) Backend {
	//nolint:gocritic // Prefer switch statement over if statement.
	switch in {
	case "influxdb":
		return BackendInfluxDB
	}
	return -1
}

const (
	BackendInfluxDB Backend = iota
)

type Config struct {
	Backend  Backend
	InfluxDB *InfluxDBConfig

	Data Data
}

type Data struct {
	Repository string
	Status     string

	Tags map[string]string
}

type ValueMustBeSetError struct {
	value string
}

func (e *ValueMustBeSetError) Error() string {
	return fmt.Sprintf("value for %q must be set", e.value)
}

// NewFromInputs creates a new Config from the GitHub Action inputs.
func NewFromInputs(action *githubactions.Action) (*Config, error) {
	c := &Config{}

	backend, err := toValidBackend(action)
	if err != nil {
		return nil, err
	}
	c.Backend = backend

	//nolint:gocritic // Prefer switch statement over if statement.
	switch backend {
	case BackendInfluxDB:
		influxDBConfig, backendErr := toInfluxDBConfig(action)
		if backendErr != nil {
			return nil, backendErr
		}
		c.InfluxDB = influxDBConfig
	}

	kvs, err := toTagsMap(action)
	if err != nil {
		return nil, err
	}
	c.Data = Data{
		Repository: action.GetInput("repository"),
		Status:     action.GetInput("status"),
		Tags:       kvs,
	}

	return c, nil
}

func toValidBackend(action *githubactions.Action) (Backend, error) {
	backendString := action.GetInput("backend")
	backend := StringToBackend(backendString)
	if backend == -1 {
		//nolint:goerr113 // No need to return a custom error.
		return -1, fmt.Errorf("unknown backend %q", backendString)
	}

	return backend, nil
}

func toInfluxDBConfig(action *githubactions.Action) (*InfluxDBConfig, error) {
	token := action.GetInput("influxdb_token")
	if token == "" {
		return nil, &ValueMustBeSetError{"influxdb_token"}
	}
	url := action.GetInput("influxdb_url")
	if url == "" {
		return nil, &ValueMustBeSetError{"influxdb_url"}
	}
	org := action.GetInput("influxdb_org")
	if org == "" {
		return nil, &ValueMustBeSetError{"influxdb_org"}
	}
	bucket := action.GetInput("influxdb_bucket")
	if bucket == "" {
		return nil, &ValueMustBeSetError{"influxdb_bucket"}
	}

	return &InfluxDBConfig{
		Token:  token,
		URL:    url,
		Org:    org,
		Bucket: bucket,
	}, nil
}

func toTagsMap(action *githubactions.Action) (map[string]string, error) {
	tagsString := action.GetInput("tags")
	// spit tags on a comma
	tags := strings.Split(tagsString, ",")
	kvs := make(map[string]string)
	for _, kv := range tags {
		// split on an eqauls sign
		kv := strings.Split(kv, "=")
		//nolint:gomnd // No need to use a constant.
		if len(kv) != 2 {
			//nolint:goerr113 // No need to return a custom error.
			return nil, fmt.Errorf("invalid tag format: %s", kv)
		}
		// trim spaces
		k := strings.TrimSpace(kv[0])
		v := strings.TrimSpace(kv[1])
		// check if tag is already set
		if _, ok := kvs[k]; ok {
			//nolint:goerr113 // No need to return a custom error.
			return nil, fmt.Errorf("tag %s is already set", k)
		}
		kvs[k] = v
	}

	return kvs, nil
}
