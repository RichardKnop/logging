package logging

// DefaultFormatter adds filename and line number before the log message
type DefaultFormatter struct {
}

// GetPrefix returns ""
func (f *DefaultFormatter) GetPrefix(lvl Level) string {
	return ""
}

// GetSuffix returns ""
func (f *DefaultFormatter) GetSuffix(lvl Level) string {
	return ""
}

// Format adds filename and line number before the log message
func (f *DefaultFormatter) Format(lvl Level, v ...interface{}) []interface{} {
	return append([]interface{}{header()}, v...)
}
