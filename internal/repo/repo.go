package repo

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"thelastcookies/api-base/internal/model"
)

type Repository interface {
	// Portlet
	GetPortlets(ctx context.Context) ([]*model.Portlet, error)
	GetPortlet(ctx context.Context, id string) (*model.Portlet, error)
	GetPortletsByIds(ctx context.Context, ids *[]string) ([]*model.Portlet, error)
	GetPortletByPortletId(ctx context.Context, portletId string) (*model.Portlet, error)
	CreatePortlet(ctx context.Context, portlet *model.Portlet) (string, error)
	UpdatePortlet(ctx context.Context, portlet *model.Portlet) error
	DeletePortlet(ctx context.Context, id string) error
	PortletIsExist(ctx context.Context, portlet *model.Portlet) (bool, error)
	// Role
	GetRolePortlet(ctx context.Context, id string) (*model.RolePortlet, error)
	GetRolePortletsByRoleId(ctx context.Context, roleId string) ([]*model.RolePortlet, error)
	CreateRolePortlet(ctx context.Context, rolePortlet *model.RolePortlet) (string, error)
	CreateRolePortlets(ctx context.Context, rolePortlets []*model.RolePortlet) ([]string, error)
	UpdateRolePortlet(ctx context.Context, rolePortlet *model.RolePortlet) error
	DeleteRolePortlet(ctx context.Context, id string) error
	DeleteRolePortletsByRoleId(ctx context.Context, roleId string) error
	// User
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
