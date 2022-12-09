package repo

import "github.com/eatmoreapple/openwechat"

type IRepository interface {
	SetMessageHandler(handler func(msg *openwechat.Message)) error
}
