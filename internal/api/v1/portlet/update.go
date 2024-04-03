package portlet

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"thelastcookies/api-base/internal/ecode"
	"thelastcookies/api-base/internal/model"
	"thelastcookies/api-base/internal/service"
	"thelastcookies/api-base/pkg/errno"
	"thelastcookies/api-base/pkg/response"
)

func Update(c *gin.Context) {
	var p *model.Portlet
	if err := c.BindJSON(&p); err != nil {
		response.Error(c, errno.ErrInvalidParam)
		return
	}
	err := service.Svc.Portlet().UpdatePortlet(c.Request.Context(), p)
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
