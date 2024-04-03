package role

import (
	"github.com/gin-gonic/gin"
	"thelastcookies/api-base/internal/ecode"
	"thelastcookies/api-base/internal/service"
	"thelastcookies/api-base/pkg/errno"
	"thelastcookies/api-base/pkg/response"
)

func PortletsAdd(c *gin.Context) {
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
	idList, err := service.Svc.RolePortlet().AddRolePortlets(c.Request.Context(), roleId, pIdList)
	if err != nil {
		response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		return
	}
	response.Success(c, &idList)
}
