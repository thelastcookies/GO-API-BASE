package role

import (
	"github.com/gin-gonic/gin"
	"tlc.platform/web-service/internal/ecode"
	"tlc.platform/web-service/internal/service"
	"tlc.platform/web-service/pkg/errno"
	"tlc.platform/web-service/pkg/response"
)

func PortletsDel(c *gin.Context) {
	roleId := c.Param("roleId")
	if roleId == "" {
		response.Error(c, ecode.ErrInvalidRoleId)
		return
	}
	err := service.Svc.RolePortletSvc.DeleteRolePortletsByRoleId(c, roleId)
	if err != nil {
		response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		return
	}
	response.Success(c, nil)
}
