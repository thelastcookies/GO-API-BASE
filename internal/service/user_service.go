package service

import (
	"context"
	"strconv"
	"thelastcookies/api-base/internal/model"
	"thelastcookies/api-base/internal/repo"
	"thelastcookies/api-base/pkg/snowflake"
)

type UserService interface {
	GetPortletsByUserId(ctx context.Context, userId string) ([]*model.Portlet, error)
	AddUserPortlets(ctx context.Context, userId string, pIdLIst []string) ([]string, error)
	UpdateUserPortlets(ctx context.Context, userId string, pIdLIst []string) ([]string, error)
	DeleteUserPortletsByUserId(ctx context.Context, userId string) error
}

type userService struct {
	repo repo.Repository
}

func newUserSvc(svc *service) *userService {
	return &userService{repo: svc.repo}
}

func (us *userService) GetPortletsByUserId(ctx context.Context, userId string) ([]*model.Portlet, error) {
	// 待补充：判断 userId 是否存在
	rpList, err := us.repo.GetUserPortletsByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	if len(rpList) == 0 {
		return make([]*model.Portlet, 0), nil
	}
	idList := make([]string, 0)
	for _, rp := range rpList {
		idList = append(idList, rp.PortletId)
	}
	pList, err := us.repo.GetPortletsByIds(ctx, &idList)
	if err != nil {
		return nil, err
	}
	return pList, nil
}

func (us *userService) AddUserPortlets(ctx context.Context, userId string, pIdLIst []string) ([]string, error) {
	var rpList []*model.UserPortlet
	for _, pId := range pIdLIst {
		rp := &model.UserPortlet{
			Id:        strconv.FormatInt(snowflake.IDGen.Snow(), 10),
			PortletId: pId,
			UserId:    userId,
		}
		rpList = append(rpList, rp)
	}
	return us.repo.CreateUserPortlets(ctx, rpList)
}

func (us *userService) UpdateUserPortlets(ctx context.Context, userId string, pIdLIst []string) ([]string, error) {
	if err := us.repo.DeleteUserPortletsByUserId(ctx, userId); err != nil {
		return []string{}, err
	}
	return us.AddUserPortlets(ctx, userId, pIdLIst)
}

func (us *userService) DeleteUserPortletsByUserId(ctx context.Context, userId string) error {
	return us.repo.DeleteUserPortletsByUserId(ctx, userId)
}
