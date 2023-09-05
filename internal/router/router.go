package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tlc.platform/web-service/config"
	"tlc.platform/web-service/internal/api/v1/portlet"
	"tlc.platform/web-service/internal/api/v1/role"
	"tlc.platform/web-service/internal/api/v1/user"
	"tlc.platform/web-service/internal/service"
	"tlc.platform/web-service/pkg/middleware"
	"tlc.platform/web-service/pkg/response"
)

var router *gin.Engine

func NewGinRouter() error {
	// 初始化 service
	service.Svc = service.New()

	// 初始化 gin router
	router = gin.Default()
	// 加载中间件
	// 配置跨域
	router.Use(middleware.Cors())

	// 404 Handler.
	router.NoRoute(response.RouteNotFound)
	router.NoMethod(response.RouteNotFound)

	// HealthCheck 健康检查路由
	router.GET("/health", response.HealthCheck)

	// 配置接口
	v1 := router.Group("/v1")
	{
		// portlet 基本接口
		v1.GET("/portlets", portlet.List)
		v1.GET("/portlet/:id", portlet.Get)
		v1.POST("/portlet", portlet.Add)
		v1.PUT("/portlet", portlet.Update)
		v1.DELETE("/portlet/:id", portlet.Del)

		// role_portlet 接口
		v1.GET("/role/:roleId/portlets", role.PortletList)
		v1.POST("/role/:roleId/portlets", role.PortletsAdd)
		v1.PUT("/role/:roleId/portlets", role.PortletsUpdate)
		v1.DELETE("/role/:roleId/portlets", role.PortletsDel)

		// user_portlet 接口
		v1.GET("/user/:userId/portlets", user.PortletList)
		v1.POST("/user/:userId/portlets", user.PortletsAdd)
		v1.PUT("/user/:userId/portlets", user.PortletsUpdate)
		v1.DELETE("/user/:userId/portlets", user.PortletsDel)
	}

	// 启动接口服务
	err := router.Run(config.Conf.HTTP.Addr)
	if err != nil {
		fmt.Errorf("failed to init server: %v", err)
	}
	return err
}
