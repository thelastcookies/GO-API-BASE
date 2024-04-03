package portlet

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"thelastcookies/api-base/internal/ecode"
	"thelastcookies/api-base/internal/service"
	"thelastcookies/api-base/pkg/errno"
	"thelastcookies/api-base/pkg/response"
)

func Del(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, errno.ErrInvalidParam)
		return
	}
	err := service.Svc.Portlet().DeletePortlet(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, ecode.ErrPortletNotFound) {
			response.Error(c, ecode.ErrPortletNotFound)
			return
		}
		response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		return
	}
	response.Success(c, nil)
}
