package user

import (
	"github.com/gin-gonic/gin"
	"thelastcookies/api-base/internal/ecode"
	"thelastcookies/api-base/internal/service"
	"thelastcookies/api-base/pkg/errno"
	"thelastcookies/api-base/pkg/response"
)

func PortletsDel(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		response.Error(c, ecode.ErrInvalidUserId)
		return
	}
	err := service.Svc.UserPortlet().DeleteUserPortletsByUserId(c, userId)
	if err != nil {
		response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		return
	}
	response.Success(c, nil)
}
