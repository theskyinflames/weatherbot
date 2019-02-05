package service

import (
	"os"

	"github.com/sirupsen/logrus"
)

type (
	BotLog struct {
		Log *logrus.Logger
	}
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.WarnLevel)
}

func NewBotLog() *BotLog {
	return &BotLog{
		Log: logrus.New(),
	}
}

func (bl BotLog) Critical(args ...interface{}) {
	bl.Log.Fatal(args)
}

func (bl BotLog) Criticalf(format string, args ...interface{}) {
	bl.Log.Fatalf(format, args)
}

func (bl BotLog) Debug(args ...interface{}) {
	bl.Log.Warn(args)
}
func (bl BotLog) Debugf(format string, args ...interface{}) {
	bl.Log.Warnf(format, args)
}

func (bl BotLog) Fatal(args ...interface{}) {
	bl.Log.Fatal(args)
}

func (bl BotLog) Fatalf(format string, args ...interface{}) {
	bl.Log.Fatal(format, args)
}

func (bl BotLog) Panic(args ...interface{}) {
	bl.Log.Panic(args)
}

func (bl BotLog) Panicf(format string, args ...interface{}) {
	bl.Log.Panicf(format, args)
}

func (bl BotLog) Error(args ...interface{}) {
	bl.Log.Error(args)
}

func (bl BotLog) Errorf(format string, args ...interface{}) {
	bl.Log.Errorf(format, args)
}

func (bl BotLog) Warning(args ...interface{}) {
	bl.Log.Warning(args)
}

func (bl BotLog) Warningf(format string, args ...interface{}) {
	bl.Log.Warningf(format, args)
}

func (bl BotLog) Notice(args ...interface{}) {
	bl.Log.Warn(args)
}

func (bl BotLog) Noticef(format string, args ...interface{}) {
	bl.Log.Warnf(format, args)
}

func (bl BotLog) Info(args ...interface{}) {
	bl.Log.Info(args)
}

func (bl BotLog) Infof(format string, args ...interface{}) {
	bl.Log.Infof(format, args)
}
