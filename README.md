<!--
 Copyright 2023 Dimitri Koshkin. All rights reserved.
 SPDX-License-Identifier: Apache-2.0
 -->

# GitHub Actions Monitor

Write GitHub Action job status to a remote backend. This project currently supports: [InfluxDB](https://www.influxdata.com/).

## Prerequisites

- An existing GitHub Action that you would like write job status to some remote backend.

## Usage Instructions

Copy the following code snippet into your GitHub Action workflow file, and replace the `backend` and the backend specific variables with your backend of choice.
Value `tags` is optional, but can be used to add additional metadata to the status.

### InfluxDB

```yaml
      - name: Run status-writer-action action
        # always run this step, even if previous steps fail
        if: always()
        uses: dkoshkin/status-writer-action@release-v0.1.0
        with:
          # select the backend to use
          backend: "influxdb"
          # set InfluxDB details
          influxdb_token: "${{ secrets.INFLUXDB_TOKEN }}"
          influxdb_url: "${{ secrets.INFLUXDB_URL }}"
          influxdb_org: "${{ secrets.INFLUXDB_ORG }}"
          influxdb_bucket: "${{ secrets.INFLUXDB_BUCKET }}"
          # set the status and additional metadata tags
          repository: "${{ github.repository }}"
          status: ${{ job.status }}
          tags: "workflow=${{ github.workflow }},job=${{ github.job }},ref=${{ github.ref_name }}"
```

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
