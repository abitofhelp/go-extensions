////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT  license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Package error implements helpful utility functions for error processing.
package error

import (
	"fmt"
	"os"
)

// Type OnErrorFunc defines an optional function to call after handling the error.
type OnErrorFunc func(err error)

// Func OnError determines whether an error has happened.
// If there is an error, it is logged, and an optional function can be invoked.
// On error, false is returned to the caller, otherwise, true.
func IsError(err error, onError OnErrorFunc) bool {
	if err != nil {
		// Log the error...
		LogError("", err)

		// Do additional handling of the error, as implemented by the caller.
		if onError != nil {
			onError(err)
		}

		return true
	}

	return false
}

// Implement the ILogger interface.
// Function LogError writes an optional message and a required error to stderr.
func LogError(message string, err error) {
	fmt.Fprint(os.Stderr, "ERROR: %s\n%+v\n", message, err)
}
