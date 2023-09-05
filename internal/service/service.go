package service

var Svc *Service

//type Service interface {
//	PortletS() *PortletService
//}

type Service struct {
	//portletRepo    portlet.Repository
	PortletSvc     *PortletService
	RolePortletSvc *RoleService
	UserPortletSvc *UserService
}

func New() *Service {
	//portletRepo := portlet.New(model.GDB)
	return &Service{
		//portletRepo:    portletRepo,
		PortletSvc:     NewPortletSvc(),
		RolePortletSvc: NewRoleSvc(),
		UserPortletSvc: NewUserSvc(),
	}
}

//func (s *Service) PortletS() *PortletService {
//	return NewPortletSvc(s)
//}
