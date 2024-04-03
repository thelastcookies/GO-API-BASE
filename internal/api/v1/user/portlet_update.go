package user

import (
	"github.com/gin-gonic/gin"
	"thelastcookies/api-base/internal/ecode"
	"thelastcookies/api-base/internal/service"
	"thelastcookies/api-base/pkg/errno"
	"thelastcookies/api-base/pkg/response"
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
	idList, err := service.Svc.UserPortlet().UpdateUserPortlets(c.Request.Context(), userId, pIdList)
	if err != nil {
		response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		return
	}
	response.Success(c, idList)
}
