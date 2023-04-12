<!--
 Copyright 2023 Dimitri Koshkin. All rights reserved.
 SPDX-License-Identifier: Apache-2.0
 -->

# GitHub Actions Monitor

[![checks](https://github.com/dkoshkin/status-writer-action/actions/workflows/checks.yml/badge.svg?branch=main)](https://github.com/dkoshkin/status-writer-action/actions/workflows/checks.yml)

Write GitHub Action job status to a remote backend.
This project currently supports:

- [InfluxDB](https://www.influxdata.com/)
- [Google Sheets](https://www.google.com/sheets/about/) (coming soon)
- [PostgreSQL](https://www.postgresql.org/) (coming soon)

## Prerequisites

- An existing GitHub Action that you would like write job status to some remote backend.
- One of the supported backends configured with appropriate credentials.

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
        actor: "${{ github.actor }}"
        status: "${{ job.status }}"
        tags: "workflow=${{ github.workflow }},job=${{ github.job }},ref=${{ github.ref_name }},run_number=${{ github.run_number }},run_id=${{ github.run_id }}"
```

See sample workflows in [.github/workflows/release-checks.yaml](.github/workflows/release-checks.yaml)
that use this GitHub Action as a [composite action](https://docs.github.com/en/actions/creating-actions/creating-a-composite-action).

### Versioning

This project uses [semantic versioning](https://semver.org/).
However, the `alpha` tag is used to indicate that the project is still in early development stage.
Once the project reaches `v1.0.0`, a new `v1` tag will be created and the `alpha` tag will be removed.
Tags `alpha` and `v1` will be updated to point to the latest release.
You may also use any of the released tags dirrectly by adding an `action` suffix, for example `v1.0.0-action`.

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
