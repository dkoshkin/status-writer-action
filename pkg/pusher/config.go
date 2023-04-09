// Copyright 2023 Dimitri Koshkin. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package pusher

import (
	"fmt"
	"strings"

	"github.com/sethvargo/go-githubactions"
)

type Config struct {
	Keys map[string]string
}

// NewFromInputs creates a new Config from the GitHub Action inputs.
func NewFromInputs(action *githubactions.Action) (*Config, error) {
	keysString := action.GetInput("keys")
	// spit keys on a comma
	keys := strings.Split(keysString, ",")
	kvs := make(map[string]string)
	for _, kv := range keys {
		// split on an eqauls sign
		kv := strings.Split(kv, "=")
		//nolint:gomnd // No need to use a constant.
		if len(kv) != 2 {
			//nolint:goerr113 // No need to return a custom error.
			return nil, fmt.Errorf("invalid key format: %s", kv)
		}
		// trim spaces
		k := strings.TrimSpace(kv[0])
		v := strings.TrimSpace(kv[1])
		// check if key is already set
		if _, ok := kvs[k]; ok {
			//nolint:goerr113 // No need to return a custom error.
			return nil, fmt.Errorf("key %s is already set", k)
		}
		kvs[k] = v
	}
	c := Config{
		kvs,
	}
	return &c, nil
}
