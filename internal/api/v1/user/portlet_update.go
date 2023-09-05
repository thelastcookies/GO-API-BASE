package user

import (
	"github.com/gin-gonic/gin"
	"tlc.platform/web-service/internal/ecode"
	"tlc.platform/web-service/internal/service"
	"tlc.platform/web-service/pkg/errno"
	"tlc.platform/web-service/pkg/response"
)

func PortletsUpdate(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		response.Error(c, ecode.ErrInvalidUserId)
		return
	}
	var pIdList []string
	if err := c.BindJSON(&pIdList); err != nil {
		response.Error(c, errno.ErrInvalidParam)
		return
	}
	idList, err := service.Svc.UserPortletSvc.UpdateUserPortlets(c.Request.Context(), userId, pIdList)
	if err != nil {
		response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		return
	}
	response.Success(c, idList)
}
