package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	config2 "thelastcookies/api-base/config"
	"thelastcookies/api-base/internal/model"
	"thelastcookies/api-base/internal/repo"
	"thelastcookies/api-base/internal/router"
	"thelastcookies/api-base/internal/service"
	"thelastcookies/api-base/pkg/config"
	"time"
)

// 全局时区设置
func initTimeZone() {
	var zoneOffset = 8 * 3600
	var cstZone = time.FixedZone("CST", zoneOffset) // 东八
	time.Local = cstZone
}

// 日志记录
func initLog() {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	// 读取配置文件
	c := config.New("config/local")
	if err := c.Load("app", "yaml", &config2.Conf); err != nil {
		panic(err)
	}

	// 初始化时区
	initTimeZone()
	// 初始化日志
	initLog()

	// 初始化数据库连接
	model.InitMySQL()

	// 初始化接口服务
	service.Svc = service.New(repo.New(model.GetDB()))

	// 启动 HTTP 服务
	if err := router.NewGinRouter(); err != nil {
		panic(err)
	}
}
