package service

import (
	"context"
	"strconv"
	"tlc.platform/web-service/internal/model"
	"tlc.platform/web-service/internal/repo/user"
	"tlc.platform/web-service/pkg/snowflake"
)

type UserServiceFunc interface {
	GetPortletsByUserId(ctx context.Context, userId string) ([]*model.Portlet, error)
	AddUserPortlets(ctx context.Context, userId string, pIdLIst []string) ([]string, error)
	UpdateUserPortlets(ctx context.Context, userId string, pIdLIst []string) ([]string, error)
	DeleteUserPortletsByUserId(ctx context.Context, userId string) error
}

type UserService struct {
	repo user.Repository
}

func NewUserSvc() *UserService {
	userPortletRepo := user.New(model.GDB)
	return &UserService{
		repo: userPortletRepo,
	}
}

func (rs *UserService) GetPortletsByUserId(ctx context.Context, userId string) ([]*model.Portlet, error) {
	// 待补充：判断 userId 是否存在
	rpList, err := rs.repo.GetUserPortletsByUserId(ctx, userId)
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
	pList, err := Svc.PortletSvc.repo.GetPortletsByIds(ctx, &idList)
	if err != nil {
		return nil, err
	}
	return pList, nil
}

func (rs *UserService) AddUserPortlets(ctx context.Context, userId string, pIdLIst []string) ([]string, error) {
	var rpList []*model.UserPortlet
	for _, pId := range pIdLIst {
		rp := &model.UserPortlet{
			ID:        strconv.FormatInt(snowflake.IDGen.Snow(), 10),
			PortletId: pId,
			UserId:    userId,
		}
		rpList = append(rpList, rp)
	}
	return rs.repo.CreateUserPortlets(ctx, rpList)
}

func (rs *UserService) UpdateUserPortlets(ctx context.Context, userId string, pIdLIst []string) ([]string, error) {
	if err := rs.repo.DeleteUserPortletsByUserId(ctx, userId); err != nil {
		return []string{}, err
	}
	return rs.AddUserPortlets(ctx, userId, pIdLIst)
}

func (rs *UserService) DeleteUserPortletsByUserId(ctx context.Context, userId string) error {
	return rs.repo.DeleteUserPortletsByUserId(ctx, userId)
}
