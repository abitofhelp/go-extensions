////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT  license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Package time contains helpful methods that extend the runtime's time package.
package time

import "time"

// The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC.
func Zero() time.Time {
	return time.Time{}
}
