////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Package logger provides an abstraction for logging, so it can be easily changed.
package logger

// Interface ILogger defines the functions to use for logging.
type ILogger interface {
	// Function LogError writes an optional message and a required error.
	LogError(message string, err error)
}
