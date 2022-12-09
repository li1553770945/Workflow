package wechat

import "HomeWorkGo/internal/repo"

type Service struct {
	Repo repo.IRepository
}

func NewWeChat(repo repo.IRepository) IWechat {
	return &Service{
		Repo: repo,
	}
}
