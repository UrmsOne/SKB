/*
@Author: urmsone urmsone@163.com
@Date: 2/23/21 8:39 PM
@Name: logger.go
*/
package utils

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
)

var (
	LogName = "server.log"
)

func NewLog(logPath string) logrus.FieldLogger {
	// TODO: 滚动日志,实现按天切割日志文件
	// https://github.com/natefinch/lumberjack
	// https://github.com/lestrrat-go/file-rotatelogs
	lg := logrus.New()

	lg.SetOutput(io.MultiWriter(os.Stdout,
		&lumberjack.Logger{
			Filename:   filepath.Join(logPath, LogName),
			MaxSize:    10,
			MaxAge:     3,
			MaxBackups: 3,
			LocalTime:  false,
			Compress:   true,
		}))

	lg.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05Z",
		FullTimestamp:   true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyFile:  "file",
			logrus.FieldKeyMsg:   "msg",
		},
		DisableColors: false,
	})

	return lg
}
