package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"thelastcookies/api-base/internal/model"
)

func (r *repository) GetRolePortlet(ctx context.Context, id string) (*model.RolePortlet, error) {
	rp := new(model.RolePortlet)
	if err := r.orm.First(&rp, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(gorm.ErrRecordNotFound, "[repo.role_portlet] record not found")
		}
		return nil, errors.Wrap(err, "[repo.role_portlet] db query err")
	}
	return rp, nil
}

func (r *repository) GetRolePortletsByRoleId(ctx context.Context, roleId string) ([]*model.RolePortlet, error) {
	rpList := make([]*model.RolePortlet, 0)
	rp := model.RolePortlet{RoleId: roleId}
	if err := r.orm.Where(&rp).Find(&rpList).Error; err != nil {
		return nil, errors.Wrap(err, "[repo.role_portlet] db query err")
	}
	return rpList, nil
}

func (r *repository) CreateRolePortlet(ctx context.Context, rolePortlet *model.RolePortlet) (string, error) {
	rp := &model.RolePortletBase{RolePortlet: *rolePortlet}
	if err := r.orm.Create(&rp).Error; err != nil {
		return "", errors.Wrap(err, "[repo.role_portlet] add record err")
	}
	return rp.Id, nil
}

func (r *repository) CreateRolePortlets(ctx context.Context, rolePortlets []*model.RolePortlet) ([]string, error) {
	var idList []string
	var rpBaseList []*model.RolePortletBase
	for _, record := range rolePortlets {
		idList = append(idList, record.Id)
		rpBaseList = append(rpBaseList, &model.RolePortletBase{RolePortlet: *record})
	}
	if err := r.orm.Create(&rpBaseList).Error; err != nil {
		return []string{}, errors.Wrap(err, "[repo.role_portlet] add record err")
	}
	return idList, nil
}

func (r *repository) UpdateRolePortlet(ctx context.Context, rolePortlet *model.RolePortlet) error {
	if err := r.orm.Model(&model.RolePortlet{}).Updates(&rolePortlet).Error; err != nil {
		return errors.Wrap(err, "[repo.role_portlet] update record err")
	}
	return nil
}

func (r *repository) DeleteRolePortlet(ctx context.Context, id string) error {
	if err := r.orm.Delete(&model.RolePortlet{Id: id}).Error; err != nil {
		return errors.Wrap(err, "[repo.role_portlet] record delete err")
	}
	return nil
}

func (r *repository) DeleteRolePortletsByRoleId(ctx context.Context, roleId string) error {
	if err := r.orm.Where("role_id = ?", roleId).Delete(&model.RolePortlet{}).Error; err != nil {
		return errors.Wrap(err, "[repo.role_portlet] record delete err")
	}
	return nil
}
