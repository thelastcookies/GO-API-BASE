package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"thelastcookies/api-base/internal/ecode"
	"thelastcookies/api-base/internal/model"
)

func (r *repository) GetPortlets(ctx context.Context) ([]*model.Portlet, error) {
	pList := make([]*model.Portlet, 0)
	if err := r.orm.Find(&pList).Error; err != nil {
		return nil, errors.Wrap(err, "[repo.portlet_base] db query err")
	}
	return pList, nil
}

func (r *repository) GetPortlet(ctx context.Context, id string) (*model.Portlet, error) {
	p := new(model.Portlet)
	if err := r.orm.First(&p, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(gorm.ErrRecordNotFound, "[repo.portlet_base] portlet not found")
		}
		return nil, errors.Wrap(err, "[repo.portlet_base] db query err")
	}
	return p, nil
}

func (r *repository) GetPortletsByIds(ctx context.Context, ids *[]string) ([]*model.Portlet, error) {
	pList := make([]*model.Portlet, 0)
	if len(*ids) == 0 {
		return nil, ecode.ErrPortletQueryConditions
	}
	if err := r.orm.Where(ids).Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{*ids}, WithoutParentheses: true},
	}).Find(&pList).Error; err != nil {
		return nil, errors.Wrap(err, "[repo.portlet_base] db query err")
	}
	return pList, nil
}

func (r *repository) GetPortletByPortletId(ctx context.Context, portletId string) (*model.Portlet, error) {
	p := model.Portlet{PortletId: portletId}
	if err := r.orm.Where(&p).First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(gorm.ErrRecordNotFound, "[repo.portlet_base] portlet not found")
		}
		return nil, errors.Wrap(err, "[repo.portlet_base] db query err")
	}
	return &p, nil
}

func (r *repository) CreatePortlet(ctx context.Context, portlet *model.Portlet) (id string, err error) {
	p := &model.PortletBase{Portlet: *portlet}
	if err := r.orm.Create(&p).Error; err != nil {
		return "", errors.Wrap(err, "[repo.portlet_base] add portlet err")
	}
	return portlet.Id, nil
}

func (r *repository) UpdatePortlet(ctx context.Context, portlet *model.Portlet) error {
	if err := r.orm.Model(&portlet).Updates(&portlet).Error; err != nil {
		return errors.Wrap(err, "[repo.portlet_base] update portlet err")
	}
	return nil
}

func (r *repository) DeletePortlet(ctx context.Context, id string) error {
	if err := r.orm.Delete(&model.Portlet{Id: id}).Error; err != nil {
		return errors.Wrap(err, "[repo.portlet_base] portlet delete err")
	}
	return nil
}

func (r *repository) PortletIsExist(ctx context.Context, portlet *model.Portlet) (bool, error) {
	err := r.orm.Where("id = ? or portlet_id = ?",
		portlet.Id, portlet.PortletId).First(&model.Portlet{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, errors.Wrap(err, "[repo.portlet_base] db query err")
	}
	return true, nil
}
