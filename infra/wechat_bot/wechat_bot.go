package wechat_bot

import (
	"context"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"log"
)

func errorHandler(err error) bool {
	fmt.Println(err)
	return true
}
func NewBot() *openwechat.Bot {
	bot := openwechat.NewBot(context.Background())

	bot.Caller.Client.SetMode(openwechat.Desktop)
	bot.Caller.Client.MaxRetryTimes = 20
	// 获取二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	// 扫码回调
	bot.ScanCallBack = func(body []byte) {
		log.Println("扫码成功,请在手机上确认登录")
	}
	// 登录回调
	bot.LoginCallBack = func(body []byte) {
		log.Println("登录成功")
	}
	bot.MessageErrorHandler = errorHandler

	// 注册登陆二维码回调
	openwechat.GetQrcodeUrl(bot.UUID())

	// 创建热存储容器对象
	reloadStorage := openwechat.NewJsonFileHotReloadStorage("storage.json")

	// 执行热登录
	err := bot.HotLogin(reloadStorage, true)
	if err != nil {
		panic("微信登陆失败:" + err.Error())
	}
	return bot
}
