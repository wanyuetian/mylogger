package mylogger

import (
	"fmt"
	"os"
	"time"
)


// FileLogger 文件日志结构体
type ConsoleLogger struct {
	level Level
}

// NewConsoleLogger 构造函数
func NewConsoleLogger(level Level) *ConsoleLogger {
	cl := &ConsoleLogger{
		level: level,
	}
	return cl
}

func (c *ConsoleLogger) log(level Level, format string, args ...interface{})  {
	logMsg := fmt.Sprintf(format, args...)
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, funcName, line := getCallerInfo(3)
	levelStr := getLevelStr(level)
	logMsg = fmt.Sprintf("[%s] [%s] [%s:%d] [%s] %s", nowStr, levelStr, fileName, line, funcName, logMsg)
	if level >= ErrorLevel{
		fmt.Fprintln(os.Stderr, logMsg)
	}

	fmt.Fprintln(os.Stdout, logMsg)
}

// Debug 记录debug日志
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	if c.level > DebugLevel {
		return
	}
	c.log(DebugLevel, format, args...)
}

// Info 记录Info日志
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	if c.level > InfoLevel {
		return
	}
	c.log(InfoLevel, format, args...)
}

// Warn 记录Warn日志
func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	if c.level > WarnLevel {
		return
	}
	c.log(WarnLevel, format, args...)
}

// Error 记录Error日志
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	if c.level > ErrorLevel {
		return
	}
	c.log(ErrorLevel, format, args...)
}

// Fatal 记录Fatal日志
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	if c.level > FatalLevel {
		return
	}
	c.log(FatalLevel, format, args...)
}

func (c *ConsoleLogger) Close(){}