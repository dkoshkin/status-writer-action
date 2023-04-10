<!--
 Copyright 2023 Dimitri Koshkin. All rights reserved.
 SPDX-License-Identifier: Apache-2.0
 -->

# GitHub Actions Monitor

Write GitHub Action job status to a remote backend. This project currently supports: [InfluxDB](https://www.influxdata.com/).

## Prerequisites

- An existing GitHub Action that you would like write job status to some remote backend.

## Usage Instructions

See sample workflows in [./.github-sample](./.github-sample)
that use this GitHub Action as a [composite action](https://docs.github.com/en/actions/creating-actions/creating-a-composite-action).

There will be a sample workflow for each supported backend.

## Setup your Dev Environment

- Install [asdf](https://asdf-vm.com/)
- Install [asdf-direnv](https://github.com/asdf-community/asdf-direnv#setup)
- Add a global `direnv` version with: `asdf global direnv latest`
- Install all tools with: `make install-tools`

Tip: to see all available make targets with descriptions, simply run `make`.

### Lint

```bash
make lint
```

### Test

```bash
make test
```

### Build

The binary for your OS will be placed in `./dist`, e.g. `./dist/status-writer-action_darwin_arm64/status-writer-action`:

```bash
make build-snapshot
```

### Pre-commit

```bash
make pre-commit
```

### Update index.js

This GitHub Action uses a javascript wrapper to call the Go binary. To update the wrapper, run:

```bash
make build-index.js
```
