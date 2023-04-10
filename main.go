// Copyright 2023 Dimitri Koshkin. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sethvargo/go-githubactions"

	"github.com/dkoshkin/status-writer-action/pkg/remote"
	"github.com/dkoshkin/status-writer-action/pkg/version"
)

func run() error {
	ctx := context.Background()
	action := githubactions.New()

	cfg, err := remote.NewFromInputs(action)
	if err != nil {
		//nolint:wrapcheck // we don't want to wrap this error
		return err
	}

	var writer remote.Writer
	//nolint:gocritic // Prefer switch statement over if statement.
	switch cfg.Backend {
	case remote.BackendInfluxDB:
		writer = remote.NewInfluxDBWriter(cfg.InfluxDB)
	}

	//nolint:wrapcheck // we don't want to wrap this error
	return writer.Write(ctx, cfg.Data)
}

func main() {
	if len(os.Args) > 1 &&
		(os.Args[1] == "version" || os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Println(version.Print())
		os.Exit(0)
	}

	err := run()
	if err != nil {
		githubactions.Fatalf("%v", err)
	}
}
