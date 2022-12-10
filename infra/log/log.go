package log

import (
	"github.com/sirupsen/logrus"
	"workflow_http/infra/conf"
)

type Log struct {
	Logger *logrus.Logger
}

func NewLog(conf *conf.Config) *Log {
	log := logrus.New()
	return &Log{
		Logger: log,
	}
}
