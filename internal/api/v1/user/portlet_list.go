package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"tlc.platform/web-service/internal/ecode"
	"tlc.platform/web-service/internal/service"
	"tlc.platform/web-service/pkg/errno"
	"tlc.platform/web-service/pkg/response"
)

func PortletList(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		response.Error(c, ecode.ErrInvalidUserId)
		return
	}
	list, err := service.Svc.UserPortletSvc.GetPortletsByUserId(c, userId)
	if err != nil {
		if errors.Is(err, ecode.ErrUserPortletsNotFound) {
			response.Success(c, &list)
		}
		response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		return
	}
	response.Success(c, &list)
}
