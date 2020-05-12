package mylogger

// 自定义一个日志库，实现日志记录的功能

// 日志级别
// DEBUG TRACE INFO WARN ERROR CRITICAL

// 日志级别
type Level uint16
const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

func getLevelStr(level Level) string{
	switch level {
	case 0:
		return "DEBUG"
	case 1:
		return "INFO"
	case 2:
		return "WARN"
	case 3:
		return "ERROR"
	case 4:
		return "CRITICAL"
	default:
		return "DEBUG"
	}
}

type Logger interface {
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Error(string, ...interface{})
	Fatal(string, ...interface{})
	Close()
}