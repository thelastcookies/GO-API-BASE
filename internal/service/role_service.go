package service

import (
	"context"
	"strconv"
	"tlc.platform/web-service/internal/model"
	"tlc.platform/web-service/internal/repo/role"
	"tlc.platform/web-service/pkg/snowflake"
)

type RoleService interface {
	GetPortletsByRoleId(ctx context.Context, roleId string) ([]*model.Portlet, error)
	AddRolePortlets(ctx context.Context, roleId string, pIdLIst []string) ([]string, error)
}

type roleService struct {
	repo role.Repository
}

func NewRoleSvc() *roleService {
	rolePortletRepo := role.New(model.GDB)
	return &roleService{
		repo: rolePortletRepo,
	}
}

func (rs *roleService) GetPortletsByRoleId(ctx context.Context, roleId string) ([]*model.Portlet, error) {
	// 待补充：判断 roleId 是否存在
	rpList, err := rs.repo.GetRolePortletsByRoleId(ctx, roleId)
	if err != nil {
		return nil, err
	}
	pList := make([]*model.Portlet, 0)
	for _, rp := range rpList {
		p, _ := Svc.PortletS().GetPortlet(ctx, rp.PortletId)
		if p != nil {
			pList = append(pList, p)
		}
	}
	return pList, nil
}

func (rs *roleService) AddRolePortlets(ctx context.Context, roleId string, pIdLIst []string) ([]string, error) {
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

func (rs *roleService) UpdateRolePortlets(ctx context.Context, roleId string, pIdLIst []string) ([]string, error) {
	if err := rs.repo.DeleteRolePortletsByRoleId(ctx, roleId); err != nil {
		return []string{}, err
	}
	return rs.AddRolePortlets(ctx, roleId, pIdLIst)
}

func (rs *roleService) DeleteRolePortletsByRoleId(ctx context.Context, roleId string) error {
	return rs.repo.DeleteRolePortletsByRoleId(ctx, roleId)
}
