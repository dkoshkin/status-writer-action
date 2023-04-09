// Copyright 2023 Dimitri Koshkin. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package pusher

import (
	"context"

	"github.com/sethvargo/go-githubactions"
)

func Run(ctx context.Context, cfg *Config) error {
	githubactions.Debugf("keys: %v", cfg.Keys)

	return nil
}
