package user

import (
	"database/sql"
	"gorm.io/gorm"
)

type Repository interface {
}

type repository struct {
	orm *gorm.DB
	db  *sql.DB
}

func New(db *gorm.DB) Repository {
	return &repository{
		orm: db,
	}
}
