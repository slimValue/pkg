package logger

// Debugf logs messages at DEBUG level.
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Infof logs messages at INFO level.
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Warnf logs messages at WARN level.
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Errorf logs messages at ERROR level.
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Fatalf logs messages at FATAL level.
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

// Debug logs messages at DEBUG level.
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Info logs messages at INFO level.
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Warn logs messages at WARN level.
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Error logs messages at ERROR level.
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Fatal logs messages at FATAL level.
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Logger is used for logging formatted messages.
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}
