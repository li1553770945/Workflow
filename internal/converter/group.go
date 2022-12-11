package converter

import (
	"workflow_http/infra/dao"
	"workflow_http/internal/domain"
)

func GroupDOToEntity(do *dao.GroupDO) *domain.GroupEntity {
	entity := &domain.GroupEntity{
		GroupNumber: do.GroupNumber,
		ID:          do.BaseModel.ID,
		Name:        do.Name,
	}
	return entity
}

func GroupMessageDOToEntity(do *dao.GroupMessageDO) *domain.GroupMessageEntity {
	entity := &domain.GroupMessageEntity{
		ID:             do.BaseModel.ID,
		Message:        do.Message,
		SenderId:       do.SenderId,
		SenderNickName: do.SenderNickName,
	}
	return entity
}

func GroupEntityToDO(entity *domain.GroupEntity) *dao.GroupDO {
	do := &dao.GroupDO{
		BaseModel:   dao.BaseModel{ID: entity.ID},
		Name:        entity.Name,
		GroupNumber: entity.GroupNumber,
		WorkflowID:  entity.WorkflowID,
	}
	return do
}

func GroupMessageEntityToDO(entity *domain.GroupMessageEntity) *dao.GroupMessageDO {
	do := &dao.GroupMessageDO{
		BaseModel:      dao.BaseModel{ID: entity.ID},
		SenderNickName: entity.SenderNickName,
		SenderId:       entity.SenderId,
		Message:        entity.Message,
	}
	return do
}
