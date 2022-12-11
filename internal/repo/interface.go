package repo

import (
	"github.com/eatmoreapple/openwechat"
	"workflow_http/internal/domain"
)

type IRepository interface {
	//workflow
	SaveWorkflow(workflow *domain.WorkflowEntity) error
	SaveWorkflowTemplate(workflowTemplate *domain.WorkflowTemplateEntity) error
	FindWorkflowTemplateByName(name string) (*domain.WorkflowTemplateEntity, error)

	//group
	SaveGroup(entity *domain.GroupEntity) error
	SaveGroupMessage(entity *domain.GroupMessageEntity) error

	SetMessageHandler(handler func(msg *openwechat.Message)) error
}
