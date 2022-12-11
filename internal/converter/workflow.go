package converter

import (
	"workflow_http/infra/dao"
	"workflow_http/internal/domain"
)

func WorkflowEntityToDO(entity *domain.WorkflowEntity) *dao.WorkflowDO {
	do := &dao.WorkflowDO{
		BaseModel:  dao.BaseModel{ID: entity.ID},
		TemplateID: entity.TemplateID,
		Content:    entity.Content,
		OpenUserId: entity.OpenUserId,
	}
	return do
}

func WorkflowTemplateEntityToDO(entity *domain.WorkflowTemplateEntity) *dao.WorkflowTemplateDO {
	do := &dao.WorkflowTemplateDO{
		BaseModel:    dao.BaseModel{ID: entity.ID},
		Name:         entity.Name,
		SolverUserId: entity.SolverUserId,
	}
	return do
}

func WorkflowTemplateDOToEntity(do *dao.WorkflowTemplateDO) *domain.WorkflowTemplateEntity {
	entity := &domain.WorkflowTemplateEntity{
		ID:           do.BaseModel.ID,
		Name:         do.Name,
		SolverUserId: do.SolverUserId,
	}
	return entity
}
