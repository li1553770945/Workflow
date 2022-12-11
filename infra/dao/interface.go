package dao

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type IDAO interface {
	CreateWorkflowTemplate(do *WorkflowTemplateDO) error
	UpdateWorkflowTemplate(do *WorkflowTemplateDO) error

	GetWorkflowTemplateByName(name string) (*WorkflowTemplateDO, error)

	CreateWorkflow(do *WorkflowDO) error
	UpdateWorkflow(do *WorkflowDO) error

	//group
	CreateGroup(do *GroupDO) error
	CreateGroupMessage(do *GroupMessageDO) error
}
