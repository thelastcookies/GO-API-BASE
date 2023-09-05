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

func Update(c *gin.Context) {
	var p *model.Portlet
	if err := c.BindJSON(&p); err != nil {
		response.Error(c, errno.ErrInvalidParam)
		return
	}
	err := service.Svc.PortletSvc.UpdatePortlet(c.Request.Context(), p)
	if err != nil {
		if errors.Is(err, ecode.ErrPortletParams) {
			response.Error(c, ecode.ErrPortletParams)
		} else if errors.Is(err, ecode.ErrPortletNotFound) {
			response.Error(c, ecode.ErrPortletNotFound)
		} else {
			response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		}
		return
	}
	response.Success(c, nil)
}
