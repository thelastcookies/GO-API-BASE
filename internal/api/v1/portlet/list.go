package portlet

import (
	"github.com/gin-gonic/gin"
	"tlc.platform/web-service/internal/service"
	"tlc.platform/web-service/pkg/errno"
	"tlc.platform/web-service/pkg/response"
)

func List(c *gin.Context) {
	list, err := service.Svc.PortletS().ListPortlet(c.Request.Context())
	if err != nil {
		response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		return
	}
	response.Success(c, &list)
}
