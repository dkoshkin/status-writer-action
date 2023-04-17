// Copyright 2023 Dimitri Koshkin. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package remote

import "time"

func timestamp() time.Time {
	return time.Now()
}

func timestampAsString() string {
	return timestamp().Format("2006-01-02T15:04:05Z")
}
