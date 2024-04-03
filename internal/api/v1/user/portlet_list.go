package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"thelastcookies/api-base/internal/ecode"
	"thelastcookies/api-base/internal/service"
	"thelastcookies/api-base/pkg/errno"
	"thelastcookies/api-base/pkg/response"
)

func PortletList(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		response.Error(c, ecode.ErrInvalidUserId)
		return
	}
	list, err := service.Svc.UserPortlet().GetPortletsByUserId(c, userId)
	if err != nil {
		if errors.Is(err, ecode.ErrUserPortletsNotFound) {
			response.Success(c, &list)
		}
		response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		return
	}
	response.Success(c, &list)
}
