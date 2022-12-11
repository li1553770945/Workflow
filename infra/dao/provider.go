package dao

import "gorm.io/gorm"

type DAO struct {
	DB *gorm.DB
}

func NewDao(DB *gorm.DB) IDAO {
	if err := DB.AutoMigrate(&WorkflowTemplateDO{}, &WorkflowDO{}, &GroupDO{}, &GroupMessageDO{}); err != nil {
		panic("failed to migrate")
	}
	return &DAO{DB: DB}
}
