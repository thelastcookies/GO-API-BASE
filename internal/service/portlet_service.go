package service

import (
	"context"
	"strconv"
	"thelastcookies/api-base/internal/ecode"
	"thelastcookies/api-base/internal/model"
	"thelastcookies/api-base/internal/repo"
	"thelastcookies/api-base/pkg/snowflake"
	"thelastcookies/api-base/pkg/utils"
)

type PortletService interface {
	ListPortlet(ctx context.Context) ([]*model.Portlet, error)
	TreePortlet(ctx context.Context) ([]*model.PortletTreeNode, error)
	GetPortlet(ctx context.Context, id string) (*model.Portlet, error)

	AddPortlet(ctx context.Context, portlet *model.Portlet) (string, error)
	UpdatePortlet(ctx context.Context, portlet *model.Portlet) error
	DeletePortlet(ctx context.Context, id string) error
}

type portletService struct {
	repo repo.Repository
}

func newPortletSvc(svc *service) *portletService {
	return &portletService{repo: svc.repo}
}

func (ps *portletService) ListPortlet(ctx context.Context) ([]*model.Portlet, error) {
	return ps.repo.GetPortlets(ctx)
}

func (ps *portletService) TreePortlet(ctx context.Context) ([]*model.PortletTreeNode, error) {
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

func (ps *portletService) GetPortlet(ctx context.Context, id string) (*model.Portlet, error) {
	return ps.repo.GetPortlet(ctx, id)
}

func (ps *portletService) AddPortlet(ctx context.Context, portlet *model.Portlet) (string, error) {
	portletId := portlet.PortletId
	if portletId == "" {
		return "", ecode.ErrInvalidPortletId
	}
	if isExist, _ := ps.repo.PortletIsExist(ctx, &model.Portlet{PortletId: portletId}); isExist {
		return "", ecode.ErrDuplicatePortletId
	}
	portlet.Id = strconv.FormatInt(snowflake.IDGen.Snow(), 10)
	return ps.repo.CreatePortlet(ctx, portlet)
}

func (ps *portletService) UpdatePortlet(ctx context.Context, portlet *model.Portlet) error {
	id := portlet.Id
	if id == "" {
		return ecode.ErrPortletParams
	}
	if isExist, _ := ps.repo.PortletIsExist(ctx, &model.Portlet{Id: id}); !isExist {
		return ecode.ErrPortletNotFound
	}
	return ps.repo.UpdatePortlet(ctx, portlet)
}

func (ps *portletService) DeletePortlet(ctx context.Context, id string) error {
	isExist, _ := ps.repo.PortletIsExist(ctx, &model.Portlet{Id: id})
	if !isExist {
		return ecode.ErrPortletNotFound
	}
	return ps.repo.DeletePortlet(ctx, id)
}
