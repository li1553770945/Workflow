package repo

import (
	"workflow_http/internal/converter"
	"workflow_http/internal/domain"
)

func (r *Repository) SaveWorkflow(workflow *domain.WorkflowEntity) error {
	if workflow.ID > 0 {
		return r.Dao.UpdateWorkflow(converter.WorkflowEntityToDO(workflow))
	} else {
		do := converter.WorkflowEntityToDO(workflow)
		err := r.Dao.CreateWorkflow(do)
		if err != nil {
			return err
		}
		workflow.ID = do.BaseModel.ID
		return nil
	}
}

func (r *Repository) SaveWorkflowTemplate(entity *domain.WorkflowTemplateEntity) error {
	if entity.ID > 0 {
		return r.Dao.UpdateWorkflowTemplate(converter.WorkflowTemplateEntityToDO(entity))
	} else {
		do := converter.WorkflowTemplateEntityToDO(entity)
		err := r.Dao.CreateWorkflowTemplate(do)
		if err != nil {
			return err
		}
		entity.ID = do.BaseModel.ID
		return nil
	}
}
func (r *Repository) FindWorkflowTemplateByName(name string) (*domain.WorkflowTemplateEntity, error) {
	do, err := r.Dao.GetWorkflowTemplateByName(name)
	if err != nil {
		return nil, err
	}
	return converter.WorkflowTemplateDOToEntity(do), err
}
