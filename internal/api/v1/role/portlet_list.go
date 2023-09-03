package role

import (
	"github.com/gin-gonic/gin"
	"tlc.platform/web-service/internal/ecode"
	"tlc.platform/web-service/internal/service"
	"tlc.platform/web-service/pkg/errno"
	"tlc.platform/web-service/pkg/response"
)

func PortletList(c *gin.Context) {
	roleId := c.Param("roleId")
	if roleId == "" {
		response.Error(c, ecode.ErrInvalidRoleId)
		return
	}
	list, err := service.Svc.RolePortletSvc.GetPortletsByRoleId(c, roleId)
	if err != nil {
		response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		return
	}
	response.Success(c, &list)
}
