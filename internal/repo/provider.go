package repo

import (
	"github.com/eatmoreapple/openwechat"
	"gorm.io/gorm"
	"workflow_http/infra/log"
)

type Repository struct {
	Db  *gorm.DB
	Bot *openwechat.Bot
	Log *log.Log
}

func NewRepository(db *gorm.DB, bot *openwechat.Bot, log *log.Log) IRepository {
	return &Repository{
		Db:  db,
		Bot: bot,
		Log: log,
	}
}
