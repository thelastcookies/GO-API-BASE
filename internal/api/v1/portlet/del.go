package portlet

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"tlc.platform/web-service/internal/ecode"
	"tlc.platform/web-service/internal/service"
	"tlc.platform/web-service/pkg/errno"
	"tlc.platform/web-service/pkg/response"
)

func Del(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, errno.ErrInvalidParam)
		return
	}
	err := service.Svc.PortletS().DeletePortlet(c.Request.Context(), id)
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
