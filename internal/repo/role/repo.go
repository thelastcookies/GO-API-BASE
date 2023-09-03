package role

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"tlc.platform/web-service/internal/model"
)

type Repository interface {
	GetRolePortlet(ctx context.Context, id string) (*model.RolePortlet, error)
	GetRolePortletsByRoleId(ctx context.Context, roleId string) ([]*model.RolePortlet, error)

	CreateRolePortlet(ctx context.Context, rolePortlet *model.RolePortlet) (string, error)
	CreateRolePortlets(ctx context.Context, rolePortlets []*model.RolePortlet) ([]string, error)

	UpdateRolePortlet(ctx context.Context, rolePortlet *model.RolePortlet) error

	DeleteRolePortlet(ctx context.Context, id string) error
	DeleteRolePortletsByRoleId(ctx context.Context, roleId string) error
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
