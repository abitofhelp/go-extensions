////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a BSD-style  license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Package error implements helpful utility functions for error processing.
package error

import (
	"fmt"
	"os"
)

// Func Must determines whether an error has happened.
// It returns the err value.
func Must(err error) error {
	if err != nil {
		LogError("", err)
	}

	return err
}

// Implement the ILogger interface.
// Function LogError writes an optional message and a required error to stderr.
func LogError(message string, err error) {
	fmt.Fprint(os.Stderr, "ERROR: %s\n%+v\n", message, err)
}
