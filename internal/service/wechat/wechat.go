package wechat

import (
	"github.com/eatmoreapple/openwechat"
	"workflow_http/internal/constant"
)

func (s *Service) MessageHandler(msg *openwechat.Message) {
	if msg.IsFriendAdd() {
		s.HandleNewFriendMessage(msg)
	} else if msg.IsSendByGroup() {
		s.HandleGroupMessage(msg)
	} else if msg.IsSendByFriend() {
		s.HandlePrivateMessage(msg)
	}

}

func (s *Service) HandleGroupMessage(msg *openwechat.Message) {

}
func (s *Service) HandlePrivateMessage(msg *openwechat.Message) {
	if msg.IsText() && msg.Content == "ping" {
		msg.ReplyText("pong")
	}
}
func (s *Service) HandleNewFriendMessage(msg *openwechat.Message) {
	friend, err := msg.Agree()
	if err != nil {
		s.Log.Logger.Error("接受好友请求错误：" + err.Error())
		return
	}
	curUser, err := msg.Bot.GetCurrentUser()
	if err != nil {
		s.Log.Logger.Println("查询当前登陆用户失败：" + err.Error())
		return
	}
	_, err = curUser.SendTextToFriend(friend, constant.FriendGreet)
	if err != nil {
		s.Log.Logger.Println("发送消息失败：" + err.Error())
		return
	}
}
