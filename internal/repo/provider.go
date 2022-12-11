package repo

import (
	"github.com/eatmoreapple/openwechat"
	"workflow_http/infra/dao"
	"workflow_http/infra/log"
)

type Repository struct {
	Bot *openwechat.Bot
	Log *log.Log
	Dao dao.IDAO
}

func NewRepository(dao dao.IDAO, bot *openwechat.Bot, log *log.Log) IRepository {
	return &Repository{
		Bot: bot,
		Log: log,
		Dao: dao,
	}
}
