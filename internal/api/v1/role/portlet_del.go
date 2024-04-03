package role

import (
	"github.com/gin-gonic/gin"
	"thelastcookies/api-base/internal/ecode"
	"thelastcookies/api-base/internal/service"
	"thelastcookies/api-base/pkg/errno"
	"thelastcookies/api-base/pkg/response"
)

func PortletsDel(c *gin.Context) {
	roleId := c.Param("roleId")
	if roleId == "" {
		response.Error(c, ecode.ErrInvalidRoleId)
		return
	}
	err := service.Svc.RolePortlet().DeleteRolePortletsByRoleId(c, roleId)
	if err != nil {
		response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		return
	}
	response.Success(c, nil)
}
