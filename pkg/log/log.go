package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

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
	}
	for k, v := range logLevelMap {
		hook, err := NewLevelHook(k, v)
		if err != nil {
			return err
		}
		logrus.AddHook(hook)
	}

	   file, err := os.OpenFile("./log/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	   if err != nil {
	       return err
	   }
	   logrus.SetOutput(file)
	return nil
}

type LevelHook struct {
	file   *os.File
	level  logrus.Level
	mu     sync.Mutex
	logger *logrus.Logger
}

func NewLevelHook(level logrus.Level, file string) (*LevelHook, error) {
	h := LevelHook{}
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	h.file = f
	h.level = level
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
	if _, err := h.file.Write(serialized); err != nil {
		return err
	}
	return nil
}
