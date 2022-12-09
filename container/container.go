package container

import (
	"HomeWorkGo/infra/conf"
	"HomeWorkGo/internal/service/wechat"
	"HomeWorkGo/internal/service/workflow"
)

type Container struct {
	WechatService   wechat.IWechat
	WorkflowService workflow.IWorkflow
	Config          *conf.Config
}

func NewContainer(wechatService wechat.IWechat, workflowService workflow.IWorkflow, config *conf.Config) *Container {
	return &Container{
		WorkflowService: workflowService,
		WechatService:   wechatService,
		Config:          config,
	}

}
