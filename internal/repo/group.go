package repo

import (
	"workflow_http/internal/converter"
	"workflow_http/internal/domain"
)

func (r *Repository) SaveGroup(entity *domain.GroupEntity) error {
	return r.Dao.CreateGroup(converter.GroupEntityToDO(entity))
}
func (r *Repository) SaveGroupMessage(entity *domain.GroupMessageEntity) error {
	return r.Dao.CreateGroupMessage(converter.GroupMessageEntityToDO(entity))
}
