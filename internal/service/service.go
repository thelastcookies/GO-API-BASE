package service

import (
	"thelastcookies/api-base/internal/repo"
)

var Svc Service

type Service interface {
	Portlet() PortletService
	RolePortlet() RoleService
	UserPortlet() UserService
}

type service struct {
	repo repo.Repository
}

func New(repo repo.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Portlet() PortletService {
	return newPortletSvc(s)
}

func (s *service) RolePortlet() RoleService {
	return newRoleSvc(s)
}

func (s *service) UserPortlet() UserService {
	return newUserSvc(s)
}
