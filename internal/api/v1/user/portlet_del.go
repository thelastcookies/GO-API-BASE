package user

import (
	"github.com/gin-gonic/gin"
	"tlc.platform/web-service/internal/ecode"
	"tlc.platform/web-service/internal/service"
	"tlc.platform/web-service/pkg/errno"
	"tlc.platform/web-service/pkg/response"
)

func PortletsDel(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		response.Error(c, ecode.ErrInvalidUserId)
		return
	}
	err := service.Svc.UserPortletSvc.DeleteUserPortletsByUserId(c, userId)
	if err != nil {
		response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		return
	}
	response.Success(c, nil)
}
