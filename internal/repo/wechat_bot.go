package repo

import (
	"github.com/eatmoreapple/openwechat"
)

func (r *Repository) SetMessageHandler(handler func(msg *openwechat.Message)) error {
	r.Bot.MessageHandler = handler
	return nil
}
