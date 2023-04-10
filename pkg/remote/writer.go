// Copyright 2023 Dimitri Koshkin. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package remote

import (
	"context"
)

type Writer interface {
	Write(ctx context.Context, data Data) error
}
