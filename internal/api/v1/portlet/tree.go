package portlet

import (
	"github.com/gin-gonic/gin"
	"thelastcookies/api-base/internal/service"
	"thelastcookies/api-base/pkg/errno"
	"thelastcookies/api-base/pkg/response"
)

func Tree(c *gin.Context) {
	tree, err := service.Svc.Portlet().TreePortlet(c.Request.Context())
	if err != nil {
		response.Error(c, errno.ErrInternalServer.WithDetails(err.Error()))
		return
	}
	response.Success(c, &tree)
}
