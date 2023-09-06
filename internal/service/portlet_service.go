package service

import (
	"context"
	"strconv"
	"tlc.platform/web-service/internal/ecode"
	"tlc.platform/web-service/internal/model"
	"tlc.platform/web-service/internal/repo/portlet"
	"tlc.platform/web-service/pkg/snowflake"
	"tlc.platform/web-service/pkg/utils"
)

type PortletServiceFunc interface {
	ListPortlet(ctx context.Context) ([]*model.Portlet, error)
	TreePortlet(ctx context.Context) ([]*model.PortletTreeNode, error)
	GetPortlet(ctx context.Context, id string) (*model.Portlet, error)

	AddPortlet(ctx context.Context, portlet *model.Portlet) (string, error)
	UpdatePortlet(ctx context.Context, portlet *model.Portlet) error
	DeletePortlet(ctx context.Context, id string) error
}

type PortletService struct {
	repo portlet.Repository
}

//func NewPortletSvc(svc *Service) *PortletService {
//	return &PortletService{repo: svc.portletRepo}
//}

func NewPortletSvc() *PortletService {
	portletRepo := portlet.New(model.GDB)
	return &PortletService{
		repo: portletRepo,
	}
}

func (ps *PortletService) ListPortlet(ctx context.Context) ([]*model.Portlet, error) {
	return ps.repo.GetPortlets(ctx)
}

func (ps *PortletService) TreePortlet(ctx context.Context) ([]*model.PortletTreeNode, error) {
	portlets, err := ps.repo.GetPortlets(ctx)
	if err != nil {
		return nil, err
	}
	var rawTreeNode []*model.PortletTreeNode
	for _, p := range portlets {
		node := &model.PortletTreeNode{Portlet: *p}
		rawTreeNode = append(rawTreeNode, node)
	}
	portletTree := utils.ListToTree(rawTreeNode, "")
	return portletTree, nil
}

func (ps *PortletService) GetPortlet(ctx context.Context, id string) (*model.Portlet, error) {
	return ps.repo.GetPortlet(ctx, id)
}

func (ps *PortletService) AddPortlet(ctx context.Context, portlet *model.Portlet) (string, error) {
	portletId := portlet.PortletId
	if portletId == "" {
		return "", ecode.ErrInvalidPortletId
	}
	if isExist, _ := ps.repo.PortletIsExist(ctx, &model.Portlet{PortletId: portletId}); isExist {
		return "", ecode.ErrDuplicatePortletId
	}
	portlet.ID = strconv.FormatInt(snowflake.IDGen.Snow(), 10)
	return ps.repo.CreatePortlet(ctx, portlet)
}

func (ps *PortletService) UpdatePortlet(ctx context.Context, portlet *model.Portlet) error {
	id := portlet.ID
	if id == "" {
		return ecode.ErrPortletParams
	}
	if isExist, _ := ps.repo.PortletIsExist(ctx, &model.Portlet{ID: id}); !isExist {
		return ecode.ErrPortletNotFound
	}
	return ps.repo.UpdatePortlet(ctx, portlet)
}

func (ps *PortletService) DeletePortlet(ctx context.Context, id string) error {
	isExist, _ := ps.repo.PortletIsExist(ctx, &model.Portlet{ID: id})
	if !isExist {
		return ecode.ErrPortletNotFound
	}
	return ps.repo.DeletePortlet(ctx, id)
}
