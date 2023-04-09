<!--
 Copyright 2023 Dimitri Koshkin. All rights reserved.
 SPDX-License-Identifier: Apache-2.0
 -->

# <PROJECT_NAME>

Replace `<PROJECT_NAME>` with the name of the project.

## Prerequisites

...

## Usage Instructions

...

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

The binary for your OS will be placed in `./dist`, e.g. `./dist/<PROJECT_NAME>_darwin_arm64/<PROJECT_NAME>`:

```bash
make build-snapshot
```

### Pre-commit

```bash
make pre-commit
```
