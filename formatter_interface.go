package logging

import (
	"fmt"
	"path/filepath"
	"runtime"
)

const (
	// Runtime caller depth
	depth = 3
)

// Formatter interface
type Formatter interface {
	GetPrefix(lvl Level) string
	Format(lvl Level, v ...interface{}) []interface{}
	GetSuffix(lvl Level) string
}

// Returns header including filename and line number
func header() string {
	_, fn, line, ok := runtime.Caller(depth)
	if !ok {
		fn = "???"
		line = 1
	}

	return fmt.Sprintf("%s:%d ", filepath.Base(fn), line)
}
