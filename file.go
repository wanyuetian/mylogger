package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)


// FileLogger 文件日志结构体
type FileLogger struct {
	level Level
	fileName string
	filePath string
	file *os.File
	errFile *os.File
	maxSize int64
}

// NewFileLogger 构造函数
func NewFileLogger(level Level, fileName, filePath string) *FileLogger {
	fl := &FileLogger{
		fileName: fileName,
		filePath: filePath,
		level: level,
		maxSize: 10*1024*1024,  // 10M
	}
	fl.initFile()
	return fl
}

func (f *FileLogger) initFile() {
	logName := path.Join(f.filePath,f.fileName)
	file, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("Open log file [%s] failed", logName))
	}
	f.file = file

	errLogName := fmt.Sprintf("%s.err", logName)
	errFile, err := os.OpenFile(errLogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("open log file %v failed", errLogName))
	}
	f.errFile = errFile
}

func (f *FileLogger) log(level Level, format string, args ...interface{})  {
	logMsg := fmt.Sprintf(format, args...)
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, funcName, line := getCallerInfo(3)
	levelStr := getLevelStr(level)
	logMsg = fmt.Sprintf("[%s] [%s] [%s:%d] [%s] %s", nowStr, levelStr, fileName, line, funcName, logMsg)

	f.file = f.splitLogFile(f.file)

	if level >= ErrorLevel{
		f.errFile = f.splitLogFile(f.errFile)
		fmt.Fprintln(f.errFile, logMsg)
	}

	fmt.Fprintln(f.file, logMsg)
}

// 封装一个切割日志的方法
func (f *FileLogger) splitLogFile(file *os.File) *os.File{
	// 切割日志
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	if fileSize > f.maxSize {
		fileName := file.Name()
		nowStr := time.Now().Format("2006-01-02 15:04:05.000")
		backupName := fmt.Sprintf("%s_%s.bak", fileName, nowStr)
		file.Close()
		os.Rename(fileName, backupName)

		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(fmt.Errorf("Open log file failed\n"))
		}
		return file
	}
	return file
}

// Debug 记录debug日志
func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > DebugLevel {
		return
	}
	f.log(DebugLevel, format, args...)
}

// Info 记录Info日志
func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > InfoLevel {
		return
	}
	f.log(InfoLevel, format, args...)
}

// Warn 记录Warn日志
func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > WarnLevel {
		return
	}
	f.log(WarnLevel, format, args...)
}

// Error 记录Error日志
func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > ErrorLevel {
		return
	}
	f.log(ErrorLevel, format, args...)
}

// Fatal 记录Fatal日志
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > FatalLevel {
		return
	}
	f.log(FatalLevel, format, args...)
}

func (f *FileLogger) Close(){
	f.file.Close()
	f.errFile.Close()
}