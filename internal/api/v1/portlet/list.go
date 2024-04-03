package portlet

import (
	"github.com/gin-gonic/gin"
	"thelastcookies/api-base/internal/service"
	"thelastcookies/api-base/pkg/errno"
	"thelastcookies/api-base/pkg/response"
)

func List(c *gin.Context) {
	list, err := service.Svc.Portlet().ListPortlet(c.Request.Context())
	if err != nil {
		response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		return
	}
	response.Success(c, &list)
}
