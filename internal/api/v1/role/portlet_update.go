package role

import (
	"github.com/gin-gonic/gin"
	"tlc.platform/web-service/internal/ecode"
	"tlc.platform/web-service/internal/service"
	"tlc.platform/web-service/pkg/errno"
	"tlc.platform/web-service/pkg/response"
)

func PortletsUpdate(c *gin.Context) {
	roleId := c.Param("roleId")
	if roleId == "" {
		response.Error(c, ecode.ErrInvalidRoleId)
		return
	}
	var pIdList []string
	if err := c.BindJSON(&pIdList); err != nil {
		response.Error(c, errno.ErrInvalidParam)
		return
	}
	idList, err := service.Svc.RolePortletSvc.UpdateRolePortlets(c.Request.Context(), roleId, pIdList)
	if err != nil {
		response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		return
	}
	response.Success(c, idList)
}
