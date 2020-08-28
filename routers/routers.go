package routers

import (
	"github.com/gin-gonic/gin"
	v1 "goworker/controllers/api/v1"
	"goworker/core/config"
	"net/http"
)

func RouterInit() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 注册图片静态资源目录
	r.StaticFS("/storage/images", http.Dir(config.AppConfig.ImageSavePath))
	r.StaticFS("/storage/files", http.Dir(config.AppConfig.FileSavePath))

	apiv1 := r.Group("api/v1")
	apiv1.GET("/hello", v1.SayHello)

	return r
}
