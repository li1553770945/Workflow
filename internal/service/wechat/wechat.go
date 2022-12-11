package wechat

import (
	"errors"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"strings"
	"workflow_http/internal/constant"
	"workflow_http/internal/domain"
	"workflow_http/internal/service"
)

func (s *Service) MessageHandler(msg *openwechat.Message) {
	fmt.Println(msg.IsText(), msg.IsSendByFriend(), msg.IsSendByGroup(), strings.HasPrefix(msg.FromUserName, "@"), !msg.IsSendBySelf())
	if msg.IsFriendAdd() {
		s.HandleNewFriendMessage(msg)
	} else if msg.IsSendByGroup() {
		s.HandleGroupMessage(msg)
	} else if msg.IsSendByFriend() {
		s.HandlePrivateMessage(msg)
	}
	fmt.Println("return")

}
func (s *Service) replyMsg(msg *openwechat.Message, content string) {
	_, err := msg.ReplyText(content)
	if err != nil {
		s.Log.Logger.Error("发送消息失败：" + err.Error())
	}
}
func (s *Service) HandleGroupMessage(msg *openwechat.Message) {

}
func (s *Service) findFriendById(bot *openwechat.Bot, id string) (*openwechat.Friend, error) {
	curUser, err := bot.GetCurrentUser()
	if err != nil {
		s.Log.Logger.Println("查询当前登陆用户失败：" + err.Error())
		return nil, err
	}
	friends, err := curUser.Friends(true)
	if err != nil {
		s.Log.Logger.Println("查询好友失败：" + err.Error())
		return nil, err
	}
	friend := friends.Search(1, func(friend *openwechat.Friend) bool { return friend.ID() == id })

	if friend.Count() == 0 {
		return nil, errors.New("无法找到resolver")
	}
	return friend[0], nil

}
func (s *Service) HandlePrivateId(msg *openwechat.Message) {
	sender, err := msg.Sender()
	if err != nil {
		s.replyMsg(msg, err.Error())
		return
	}
	s.replyMsg(msg, sender.ID())
}
func (s *Service) HandlePrivateCommand(msg *openwechat.Message) {
	switch msg.Content {
	case "/id":
		{
			s.HandlePrivateId(msg)
			return
		}
	default:
		s.replyMsg(msg, constant.NoSuchCommand)
	}

}
func (s *Service) HandlePrivateMessage(msg *openwechat.Message) {
	if msg.IsText() {
		if msg.Content[0] == '/' {
			s.HandlePrivateCommand(msg)
			return
		}

		workflowName, workflowContent, err := service.SplitMessage(msg.Content)
		if err != nil {
			s.replyMsg(msg, err.Error())
			return
		}

		//查询template
		workflowTemplate, err := s.Repo.FindWorkflowTemplateByName(workflowName)
		if err != nil {
			s.replyMsg(msg, err.Error())
			return
		}
		sender, err := msg.Sender()
		if err != nil {
			s.replyMsg(msg, err.Error())
			return
		}
		//创建workflow
		workflow := &domain.WorkflowEntity{
			TemplateID: workflowTemplate.ID,
			Template:   workflowTemplate,
			Content:    workflowContent,
			OpenUserId: sender.ID(),
		}
		err = s.Repo.SaveWorkflow(workflow)
		if err != nil {
			s.replyMsg(msg, "创建workflow失败："+err.Error())
			return
		}

		groupName := workflowContent
		groupEntity := &domain.GroupEntity{
			GroupNumber: "123",
			WorkflowID:  workflow.ID,
			Name:        groupName,
		}
		err = s.Repo.SaveGroup(groupEntity)
		if err != nil {
			s.replyMsg(msg, "创建群聊失败："+err.Error())
			return
		}
		curUser, err := msg.Bot.GetCurrentUser()
		if err != nil {
			s.Log.Logger.Println("查询当前登陆用户失败：" + err.Error())
			return
		}
		resolver, err := s.findFriendById(msg.Bot, workflowTemplate.SolverUserId)
		if err != nil {
			s.replyMsg(msg, "创建群聊失败:"+err.Error())
			return
		}
		fmt.Println(sender.NickName, resolver.NickName)
		group, err := curUser.CreateGroup(groupName, &openwechat.Friend{User: sender}, resolver)
		if err != nil {
			s.replyMsg(msg, "创建群聊失败:"+err.Error())
			return
		}
		_, err = group.SendText(constant.GroupGreetPrefix + workflowContent + constant.GroupGreetSuffix)
		if err != nil {
			s.replyMsg(msg, "创建群聊失败:"+err.Error())
			return
		}

	} else if msg.IsPicture() || msg.IsEmoticon() {
		s.replyMsg(msg, constant.UnSupportMessage)
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
