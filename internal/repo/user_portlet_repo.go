package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"thelastcookies/api-base/internal/model"
)

func (r *repository) GetUserPortlet(ctx context.Context, id string) (*model.UserPortlet, error) {
	rp := new(model.UserPortlet)
	if err := r.orm.First(&rp, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(gorm.ErrRecordNotFound, "[repo.user_portlet] record not found")
		}
		return nil, errors.Wrap(err, "[repo.user_portlet] db query err")
	}
	return rp, nil
}

func (r *repository) GetUserPortletsByUserId(ctx context.Context, userId string) ([]*model.UserPortlet, error) {
	rpList := make([]*model.UserPortlet, 0)
	rp := model.UserPortlet{UserId: userId}
	if err := r.orm.Where(&rp).Find(&rpList).Error; err != nil {
		return nil, errors.Wrap(err, "[repo.user_portlet] db query err")
	}
	return rpList, nil
}

func (r *repository) CreateUserPortlet(ctx context.Context, userPortlet *model.UserPortlet) (string, error) {
	rp := &model.UserPortletBase{UserPortlet: *userPortlet}
	if err := r.orm.Create(&rp).Error; err != nil {
		return "", errors.Wrap(err, "[repo.user_portlet] add record err")
	}
	return rp.Id, nil
}

func (r *repository) CreateUserPortlets(ctx context.Context, userPortlets []*model.UserPortlet) ([]string, error) {
	var idList []string
	var rpBaseList []*model.UserPortletBase
	for _, record := range userPortlets {
		idList = append(idList, record.Id)
		rpBaseList = append(rpBaseList, &model.UserPortletBase{UserPortlet: *record})
	}
	if err := r.orm.Create(&rpBaseList).Error; err != nil {
		return []string{}, errors.Wrap(err, "[repo.user_portlet] add record err")
	}
	return idList, nil
}

func (r *repository) UpdateUserPortlet(ctx context.Context, userPortlet *model.UserPortlet) error {
	if err := r.orm.Model(&model.UserPortlet{}).Updates(&userPortlet).Error; err != nil {
		return errors.Wrap(err, "[repo.user_portlet] update record err")
	}
	return nil
}

func (r *repository) DeleteUserPortlet(ctx context.Context, id string) error {
	if err := r.orm.Delete(&model.UserPortlet{Id: id}).Error; err != nil {
		return errors.Wrap(err, "[repo.user_portlet] record delete err")
	}
	return nil
}

func (r *repository) DeleteUserPortletsByUserId(ctx context.Context, userId string) error {
	if err := r.orm.Where("user_id = ?", userId).Delete(&model.UserPortlet{}).Error; err != nil {
		return errors.Wrap(err, "[repo.user_portlet] record delete err")
	}
	return nil
}
