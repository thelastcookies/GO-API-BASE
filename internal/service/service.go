package service

import (
	"tlc.platform/web-service/internal/model"
	"tlc.platform/web-service/internal/repo/portlet"
)

var Svc *service

type Service interface {
	PortletS() *portletService
	//RolePortletS() *roleService
}

type service struct {
	portletRepo    portlet.Repository
	RolePortletSvc *roleService
}

func New() *service {
	portletRepo := portlet.New(model.GDB)
	return &service{
		portletRepo:    portletRepo,
		RolePortletSvc: NewRoleSvc(),
	}
}

func (s *service) PortletS() *portletService {
	return NewPortletSvc(s)
}
