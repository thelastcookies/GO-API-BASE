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

func Add(c *gin.Context) {
	var p *model.Portlet
	if err := c.BindJSON(&p); err != nil {
		response.Error(c, errno.ErrInvalidParam)
		return
	}
	id, err := service.Svc.Portlet().AddPortlet(c.Request.Context(), p)
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
