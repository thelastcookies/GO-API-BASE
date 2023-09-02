package service

import (
	"tlc.platform/web-service/internal/model"
	"tlc.platform/web-service/internal/repo/portlet"
)

var Svc Service

type Service interface {
	PortletS() *portletService
}

type service struct {
	portletRepo portlet.Repository
}

func New() *service {
	portletRepo := portlet.New(model.GDB)
	return &service{
		portletRepo: portletRepo,
	}
}

func (s *service) PortletS() *portletService {
	return NewPortletSvc(s)
}
