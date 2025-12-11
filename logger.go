package logging

import (
	"io"
	"log"
	"os"
)

// Level type
type Level int

const (
	// DEBUG Level
	DEBUG Level = iota
	// INFO Level
	INFO
	// WARNING Level
	WARNING
	// ERROR Level
	ERROR
	// FATAL Level
	FATAL

	flag = log.Ldate | log.Ltime
)

// Log Level prefix map
var prefix = map[Level]string{
	DEBUG:   "DEBUG: ",
	INFO:    "INFO: ",
	WARNING: "WARNING: ",
	ERROR:   "ERROR: ",
	FATAL:   "FATAL: ",
}

// Logger ...
type Logger map[Level]LoggerInterface

// New returns instance of Logger.
func New(out, errOut io.Writer, ops ...ConfigOption) Logger {
	config := defaultConfig()

	for _, op := range ops {
		op(config)
	}

	// If log level is out of bounds, set to nearest valid level.
	if config.LogLevel < DEBUG {
		config.LogLevel = DEBUG
	}

	if config.LogLevel > FATAL {
		config.LogLevel = FATAL
	}

	// Fall back to stdout if out not set
	if out == nil {
		out = os.Stdout
	}

	// Fall back to stderr if errOut not set
	if errOut == nil {
		errOut = os.Stderr
	}

	l := make(map[Level]LoggerInterface, 5)

	for level := DEBUG; level <= FATAL; level++ {
		l[level] = NewNoOp()

		if level >= config.LogLevel {
			l[level] = &Wrapper{
				lvl:       level,
				formatter: config.Formatter,
				logger:    log.New(out, config.Formatter.GetPrefix(level)+prefix[level], flag),
			}
		}
	}

	return l
}

// Wrapper ...
type Wrapper struct {
	lvl       Level
	formatter Formatter
	logger    LoggerInterface
}

// Print ...
func (w *Wrapper) Print(v ...interface{}) {
	v = w.formatter.Format(w.lvl, v...)
	v = append(v, w.formatter.GetSuffix(w.lvl))
	w.logger.Print(v...)
}

// Printf ...
func (w *Wrapper) Printf(format string, v ...interface{}) {
	suffix := w.formatter.GetSuffix(w.lvl)
	v = w.formatter.Format(w.lvl, v...)
	w.logger.Printf("%s"+format+suffix, v...)
}

// Println ...
func (w *Wrapper) Println(v ...interface{}) {
	v = w.formatter.Format(w.lvl, v...)
	v = append(v, w.formatter.GetSuffix(w.lvl))
	w.logger.Println(v...)
}

// Fatal ...
func (w *Wrapper) Fatal(v ...interface{}) {
	v = w.formatter.Format(w.lvl, v...)
	v = append(v, w.formatter.GetSuffix(w.lvl))
	w.logger.Fatal(v...)
}

// Fatalf ...
func (w *Wrapper) Fatalf(format string, v ...interface{}) {
	suffix := w.formatter.GetSuffix(w.lvl)
	v = w.formatter.Format(w.lvl, v...)
	w.logger.Fatalf("%s"+format+suffix, v...)
}

// Fatalln ...
func (w *Wrapper) Fatalln(v ...interface{}) {
	v = w.formatter.Format(w.lvl, v...)
	v = append(v, w.formatter.GetSuffix(w.lvl))
	w.logger.Fatalln(v...)
}

// Panic ...
func (w *Wrapper) Panic(v ...interface{}) {
	v = w.formatter.Format(w.lvl, v...)
	v = append(v, w.formatter.GetSuffix(w.lvl))
	w.logger.Fatal(v...)
}

// Panicf ...
func (w *Wrapper) Panicf(format string, v ...interface{}) {
	suffix := w.formatter.GetSuffix(w.lvl)
	v = w.formatter.Format(w.lvl, v...)
	w.logger.Panicf("%s"+format+suffix, v...)
}

// Panicln ...
func (w *Wrapper) Panicln(v ...interface{}) {
	v = w.formatter.Format(w.lvl, v...)
	v = append(v, w.formatter.GetSuffix(w.lvl))
	w.logger.Panicln(v...)
}
