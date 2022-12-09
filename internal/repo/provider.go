package repo

import "gorm.io/gorm"

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &Repository{
		Db: db,
	}
}
