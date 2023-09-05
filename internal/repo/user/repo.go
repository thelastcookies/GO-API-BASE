package user

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"tlc.platform/web-service/internal/model"
)

type Repository interface {
	GetUserPortlet(ctx context.Context, id string) (*model.UserPortlet, error)
	GetUserPortletsByUserId(ctx context.Context, userId string) ([]*model.UserPortlet, error)

	CreateUserPortlet(ctx context.Context, userPortlet *model.UserPortlet) (string, error)
	CreateUserPortlets(ctx context.Context, userPortlets []*model.UserPortlet) ([]string, error)

	UpdateUserPortlet(ctx context.Context, userPortlet *model.UserPortlet) error

	DeleteUserPortlet(ctx context.Context, id string) error
	DeleteUserPortletsByUserId(ctx context.Context, userId string) error
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
