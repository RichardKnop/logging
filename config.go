package logging

// Config holds configuration options for the logger.
type Config struct {
	// LogLevel sets the minimum log level to output. If not set, INFO level is used.
	LogLevel Level

	// Formatter sets the log formatter to use. If not set, DefaultFormatter is used.
	Formatter Formatter
}

func defaultConfig() *Config {
	return &Config{
		LogLevel:  INFO,
		Formatter: new(DefaultFormatter),
	}
}

// ConfigOption defines a function type for configuring the logger.
type ConfigOption func(*Config)

// WithLogLevel sets the log level in the logger configuration.
func WithLogLevel(lvl Level) ConfigOption {
	return func(c *Config) {
		c.LogLevel = lvl
	}
}

// WithFormatter sets the log formatter in the logger configuration.
func WithFormatter(f Formatter) ConfigOption {
	return func(c *Config) {
		c.Formatter = f
	}
}
