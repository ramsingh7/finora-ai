package filemanager

import "errors"

var (
	// ErrInvalidPath is returned when a relative path escapes the managed root or is unsafe.
	ErrInvalidPath = errors.New("invalid path")
	// ErrTooLarge is returned when a read/write exceeds the configured max size.
	ErrTooLarge = errors.New("file too large")
	// ErrNotFound is returned when the target path does not exist under the root.
	ErrNotFound = errors.New("not found")
)
