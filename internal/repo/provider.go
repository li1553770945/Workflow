package repo

import (
	"github.com/eatmoreapple/openwechat"
	"gorm.io/gorm"
)

type Repository struct {
	Db  *gorm.DB
	Bot *openwechat.Bot
}

func NewRepository(db *gorm.DB, bot *openwechat.Bot) IRepository {
	return &Repository{
		Db:  db,
		Bot: bot,
	}
}
