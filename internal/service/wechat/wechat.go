package wechat

import "github.com/eatmoreapple/openwechat"

func MessageHandler(msg *openwechat.Message) {
	if msg.IsText() && msg.Content == "ping" {
		msg.ReplyText("pong")
	}
}
