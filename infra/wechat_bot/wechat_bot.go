package wechat_bot

import (
	"github.com/eatmoreapple/openwechat"
)

func NewBot() *openwechat.Bot {
	bot := openwechat.DefaultBot(openwechat.Desktop)

	// 注册登陆二维码回调
	openwechat.GetQrcodeUrl(bot.UUID())
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 创建热存储容器对象
	reloadStorage := openwechat.NewJsonFileHotReloadStorage("storage.json")

	// 执行热登录
	err := bot.HotLogin(reloadStorage)
	if err != nil {
		panic("微信登陆失败" + err.Error())
	}
	return bot
}
