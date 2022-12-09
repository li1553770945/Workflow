package container

import (
	"workflow_http/infra/conf"
	"workflow_http/internal/service/wechat"
	"workflow_http/internal/service/workflow"
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
