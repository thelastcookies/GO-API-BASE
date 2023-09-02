package portlet

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"tlc.platform/web-service/internal/ecode"
	"tlc.platform/web-service/internal/model"
	"tlc.platform/web-service/internal/service"
	"tlc.platform/web-service/pkg/errno"
	"tlc.platform/web-service/pkg/response"
)

func Add(c *gin.Context) {
	var p model.Portlet
	if err := c.BindJSON(&p); err != nil {
		response.Error(c, errno.ErrInvalidParam)
		return
	}
	id, err := service.Svc.PortletS().AddPortlet(c.Request.Context(), &p)
	if err != nil {
		if errors.Is(err, ecode.ErrInvalidPortletId) {
			response.Error(c, ecode.ErrInvalidPortletId)
		} else if errors.Is(err, ecode.ErrDuplicatePortletId) {
			response.Error(c, ecode.ErrDuplicatePortletId)
		} else {
			response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		}
		return
	}
	response.Success(c, id)
}
