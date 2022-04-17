package log

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sync"
)

//logrus 压测的时候发现logrus为同步日志库，性能较低 需更换为异步日志库
func Init(level logrus.Level) error {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetReportCaller(true)
	logrus.SetLevel(level)
	logLevelMap := map[logrus.Level]string{
		logrus.DebugLevel: "./log/debug.log",
		logrus.InfoLevel:  "./log/info.log",
		logrus.WarnLevel:  "./log/warn.log",
		logrus.ErrorLevel: "./log/err.log",
		logrus.PanicLevel: "./log/panic.log",
	}
	for k, v := range logLevelMap {
		hook, err := NewLevelHook(k, v)
		if err != nil {
			return err
		}
		logrus.AddHook(hook)
	}

	file, err := os.OpenFile("./log/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	buf := bufio.NewWriter(file)

	if err != nil {
		return err
	}
	logrus.SetOutput(buf)
	return nil
}

type LevelHook struct {
	file   *os.File
	level  logrus.Level
	mu     sync.Mutex
	logger io.Writer
}

func NewLevelHook(level logrus.Level, file string) (*LevelHook, error) {
	h := LevelHook{}

	logFile, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	h.level = level
	h.logger = logFile
	return &h, nil
}

func (h *LevelHook) Levels() []logrus.Level {
	return []logrus.Level{
		h.level,
	}
}

func (h *LevelHook) Fire(entry *logrus.Entry) error {
	serialized, err := entry.Logger.Formatter.Format(entry)
	if err != nil {
		return err
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	h.logger.Write(serialized)
	return nil
}
