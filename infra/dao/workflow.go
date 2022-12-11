package dao

import (
	"errors"
	"gorm.io/gorm"
)

type WorkflowTemplateDO struct {
	BaseModel BaseModel `gorm:"embedded"`

	Name         string
	SolverUserId string
}
type WorkflowDO struct {
	BaseModel BaseModel `gorm:"embedded"`

	TemplateID uint
	Template   *WorkflowTemplateDO
	Content    string
	OpenUserId string
}

func (dao *DAO) CreateWorkflowTemplate(do *WorkflowTemplateDO) error {
	err := dao.DB.Create(&do).Error
	return err
}
func (dao *DAO) UpdateWorkflowTemplate(do *WorkflowTemplateDO) error {
	err := dao.DB.Save(&do).Error
	return err
}
func (dao *DAO) GetWorkflowTemplateByName(name string) (*WorkflowTemplateDO, error) {
	workflowTemplate := &WorkflowTemplateDO{}
	err := dao.DB.Where("name = ?", name).First(workflowTemplate).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("未查询到相关workflow")
	}
	return workflowTemplate, err
}

func (dao *DAO) CreateWorkflow(do *WorkflowDO) error {
	err := dao.DB.Create(&do).Error
	return err
}

func (dao *DAO) UpdateWorkflow(do *WorkflowDO) error {
	err := dao.DB.Save(&do).Error
	return err
}
