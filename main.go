// Copyright 2023 Dimitri Koshkin. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sethvargo/go-githubactions"

	"github.com/dkoshkin/gha-monitor/pkg/pusher"
	"github.com/dkoshkin/gha-monitor/pkg/version"
)

func run() error {
	ctx := context.Background()
	action := githubactions.New()

	cfg, err := pusher.NewFromInputs(action)
	if err != nil {
		//nolint:wrapcheck // we don't want to wrap this error
		return err
	}

	//nolint:wrapcheck // we don't want to wrap this error
	return pusher.Run(ctx, cfg)
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
