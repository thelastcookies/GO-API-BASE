package portlet

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"thelastcookies/api-base/internal/ecode"
	"thelastcookies/api-base/internal/service"
	"thelastcookies/api-base/pkg/errno"
	"thelastcookies/api-base/pkg/response"
)

func Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, errno.ErrInvalidParam)
		return
	}
	p, err := service.Svc.Portlet().GetPortlet(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, ecode.ErrPortletNotFound)
			return
		}
		response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		return
	}
	response.Success(c, &p)
}
