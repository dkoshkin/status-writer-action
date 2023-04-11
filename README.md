<!--
 Copyright 2023 Dimitri Koshkin. All rights reserved.
 SPDX-License-Identifier: Apache-2.0
 -->

# GitHub Actions Monitor

Write GitHub Action job status to a remote backend.
This project currently supports:

- [InfluxDB](https://www.influxdata.com/)

## Prerequisites

- An existing GitHub Action that you would like write job status to some remote backend.

## Usage Instructions

Add the following step to your GitHub Action workflows for all jobs that you would like to monitor:

```yaml
    - name: Push job status and other metadata to a remote backend
      # always run this step, even if previous steps fail
      uses: dkoshkin/status-writer-action@alpha
      # always run this step, even if previous steps fail
      if: always()
      with:
        # select the backend to use
        backend: "influxdb"
        # set InfluxDB details
        influxdb_token: "${{ secrets.INFLUXDB_TOKEN }}"
        influxdb_url: "${{ secrets.INFLUXDB_URL }}"
        influxdb_org: "${{ secrets.INFLUXDB_ORG }}"
        influxdb_bucket: "${{ secrets.INFLUXDB_BUCKET }}"
        # set the repository, status and additional metadata tags
        repository: "${{ github.repository }}"
        status: "${{ job.status }}"
        tags: "workflow=${{ github.workflow }},job=${{ github.job }},ref=${{ github.ref_name }}"
```

See sample workflows in [.github/workflows/release-checks.yaml](.github/workflows/release-checks.yaml)
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
