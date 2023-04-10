# Copyright 2023 Dimitri Koshkin. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

.PHONY: build-index.js
build-index.js: ## Builds an index.js file
	npm i -g @vercel/ncc
	ncc build invoke-binary.js --license LICENSE -o .
