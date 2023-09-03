package portlet

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"tlc.platform/web-service/internal/model"
)

type Repository interface {
	GetPortlets(ctx context.Context) ([]*model.Portlet, error)
	GetPortlet(ctx context.Context, id string) (*model.Portlet, error)
	GetPortletsByIds(ctx context.Context, ids *[]string) ([]*model.Portlet, error)
	GetPortletByPortletId(ctx context.Context, portletId string) (*model.Portlet, error)
	CreatePortlet(ctx context.Context, portlet *model.Portlet) (string, error)
	UpdatePortlet(ctx context.Context, portlet *model.Portlet) error
	DeletePortlet(ctx context.Context, id string) error
	PortletIsExist(ctx context.Context, portlet *model.Portlet) (bool, error)
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
