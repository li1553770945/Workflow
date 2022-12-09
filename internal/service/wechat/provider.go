package wechat

import (
	"workflow_http/internal/repo"
)

type Service struct {
	Repo repo.IRepository
}

func NewWeChat(repo repo.IRepository) IWechat {
	// 注册消息处理函数
	repo.SetMessageHandler(MessageHandler)
	return &Service{
		Repo: repo,
	}
}
