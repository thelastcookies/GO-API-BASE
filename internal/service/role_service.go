package service

import (
	"context"
	"strconv"
	"tlc.platform/web-service/internal/model"
	"tlc.platform/web-service/internal/repo/role"
	"tlc.platform/web-service/pkg/snowflake"
)

type RoleServiceFunc interface {
	GetPortletsByRoleId(ctx context.Context, roleId string) ([]*model.Portlet, error)
	AddRolePortlets(ctx context.Context, roleId string, pIdLIst []string) ([]string, error)
	UpdateRolePortlets(ctx context.Context, roleId string, pIdLIst []string) ([]string, error)
	DeleteRolePortletsByRoleId(ctx context.Context, roleId string) error
}

type RoleService struct {
	repo role.Repository
}

func NewRoleSvc() *RoleService {
	rolePortletRepo := role.New(model.GDB)
	return &RoleService{
		repo: rolePortletRepo,
	}
}

func (rs *RoleService) GetPortletsByRoleId(ctx context.Context, roleId string) ([]*model.Portlet, error) {
	// 待补充：判断 roleId 是否存在
	rpList, err := rs.repo.GetRolePortletsByRoleId(ctx, roleId)
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

func (rs *RoleService) AddRolePortlets(ctx context.Context, roleId string, pIdLIst []string) ([]string, error) {
	var rpList []*model.RolePortlet
	for _, pId := range pIdLIst {
		rp := &model.RolePortlet{
			ID:        strconv.FormatInt(snowflake.IDGen.Snow(), 10),
			PortletId: pId,
			RoleId:    roleId,
		}
		rpList = append(rpList, rp)
	}
	return rs.repo.CreateRolePortlets(ctx, rpList)
}

func (rs *RoleService) UpdateRolePortlets(ctx context.Context, roleId string, pIdLIst []string) ([]string, error) {
	if err := rs.repo.DeleteRolePortletsByRoleId(ctx, roleId); err != nil {
		return []string{}, err
	}
	return rs.AddRolePortlets(ctx, roleId, pIdLIst)
}

func (rs *RoleService) DeleteRolePortletsByRoleId(ctx context.Context, roleId string) error {
	return rs.repo.DeleteRolePortletsByRoleId(ctx, roleId)
}
