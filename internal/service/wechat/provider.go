package wechat

import (
	"workflow_http/infra/log"
	"workflow_http/internal/repo"
)

type Service struct {
	Repo repo.IRepository
	Log  *log.Log
}

func NewWeChat(repo repo.IRepository, log *log.Log) IWechat {
	// 注册消息处理函数
	s := &Service{
		Repo: repo,
		Log:  log,
	}
	repo.SetMessageHandler(s.MessageHandler)
	return s
}
