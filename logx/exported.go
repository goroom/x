package logx

var (
	_defaultLogger      *Logger
	_exportedCallerSkip = 5
)

func init() {
	_defaultLogger = NewLogger()
}

func GetDefaultLogger() *Logger {
	return _defaultLogger
}

func InitDefaultLogger(opts ...Options) {
	_defaultLogger.Close()
	_defaultLogger = NewLogger(opts...)
}

func Debug(args ...interface{}) {
	if _defaultLogger.minLevel > DEBUG {
		return
	}
	_defaultLogger.log(_exportedCallerSkip, DEBUG, args)
}

func Debugf(format string, args ...interface{}) {
	if _defaultLogger.minLevel > DEBUG {
		return
	}
	_defaultLogger.logf(_exportedCallerSkip, DEBUG, format, args)
}

func Info(args ...interface{}) {
	if _defaultLogger.minLevel > INFO {
		return
	}
	_defaultLogger.log(_exportedCallerSkip, INFO, args)
}

func Infof(format string, args ...interface{}) {
	if _defaultLogger.minLevel > INFO {
		return
	}
	_defaultLogger.logf(_exportedCallerSkip, INFO, format, args)
}

func Warn(args ...interface{}) {
	if _defaultLogger.minLevel > WARN {
		return
	}
	_defaultLogger.log(_exportedCallerSkip, WARN, args)
}

func Warnf(format string, args ...interface{}) {
	if _defaultLogger.minLevel > WARN {
		return
	}
	_defaultLogger.logf(_exportedCallerSkip, WARN, format, args)
}

func Error(args ...interface{}) {
	if _defaultLogger.minLevel > ERROR {
		return
	}
	_defaultLogger.log(_exportedCallerSkip, ERROR, args)
}

func Errorf(format string, args ...interface{}) {
	if _defaultLogger.minLevel > ERROR {
		return
	}
	_defaultLogger.logf(_exportedCallerSkip, ERROR, format, args)
}

func Wait() {
	_defaultLogger.Wait()
}
