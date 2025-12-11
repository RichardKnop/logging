package logging_test

import (
	"testing"

	"github.com/RichardKnop/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name         string
		logLevel     logging.Level
		expectedLogs int
	}{
		{"DEBUG", logging.DEBUG, 5},
		{"INFO", logging.INFO, 4},
		{"WARNING", logging.WARNING, 3},
		{"ERROR", logging.ERROR, 2},
		{"FATAL", logging.FATAL, 1},
	}

	for _, tt := range tests {
		out := &stubWriter{}

		t.Run("GIVEN a logging constructor configured with min log level "+tt.name, func(t *testing.T) {
			logger := logging.New(out, nil, logging.WithLogLevel(tt.logLevel))

			t.Run("WHEN logging with all the log levels", func(t *testing.T) {
				for level := logging.DEBUG; level <= logging.FATAL; level++ {
					logger[level].Print("This is a test log")
				}

				t.Run("THEN only those allowed logs are performed", func(t *testing.T) {
					require.Len(t, out.loggedMessages, tt.expectedLogs)
					assert.Contains(t, out.loggedMessages[0], "This is a test log")
				})
			})
		})
	}
}

type stubWriter struct {
	loggedMessages []string
}

func (s *stubWriter) Write(p []byte) (n int, err error) {
	s.loggedMessages = append(s.loggedMessages, string(p))
	return len(p), nil
}
