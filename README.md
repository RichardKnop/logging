## Logging

A simple leveled logging library with coloured output.

---

Log levels:

- `INFO` (blue)
- `WARNING` (pink)
- `ERROR` (red)
- `FATAL` (red)

Formatters:

- `DefaultFormatter`
- `ColouredFormatter`

## Configuration

The logger can be configured using optional configuration functions:

- `WithLogLevel(level Level)` - Sets the minimum log level to output (default: `INFO`)
- `WithFormatter(formatter Formatter)` - Sets the log formatter (default: `DefaultFormatter`)

## Example Usage

### Basic Usage

Create a new package `log` in your app:

```go
package log

import (
	"github.com/RichardKnop/logging"
)

var (
	// Create logger with default configuration (INFO level, DefaultFormatter)
	logger = logging.New(nil, nil)

	// INFO ...
	INFO = logger[logging.INFO]
	// WARNING ...
	WARNING = logger[logging.WARNING]
	// ERROR ...
	ERROR = logger[logging.ERROR]
	// FATAL ...
	FATAL = logger[logging.FATAL]
)
```

### Custom Configuration

You can customize the logger using configuration options:

```go
package log

import (
	"github.com/RichardKnop/logging"
)

var (
	// Create logger with coloured formatter and DEBUG level
	logger = logging.New(
		nil, 
		nil, 
		logging.WithLogLevel(logging.DEBUG),
		logging.WithFormatter(new(logging.ColouredFormatter)),
	)

	// DEBUG ...
	DEBUG = logger[logging.DEBUG]
	// INFO ...
	INFO = logger[logging.INFO]
	// WARNING ...
	WARNING = logger[logging.WARNING]
	// ERROR ...
	ERROR = logger[logging.ERROR]
	// FATAL ...
	FATAL = logger[logging.FATAL]
)
```

### Using the Logger

Then from your app you could do:

```go
package main

import (
	"github.com/yourusername/yourapp/log"
)

func main() {
	log.INFO.Print("log message")
	log.WARNING.Printf("formatted %s", "message")
	log.ERROR.Println("error message")
}
```
