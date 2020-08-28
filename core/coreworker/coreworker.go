package coreworker

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goworker/core/config"
	"goworker/routers"
	"log"
	"net/http"
	_ "goworker/bootstrap"
)


/*
核心启动.
 */
func RunAll()  {
	gin.SetMode(config.ServerConfig.RunMode)
	routersInit := routers.RouterInit()
	readTimeout := config.ServerConfig.ReadTimeout
	writeTimeout := config.ServerConfig.WriteTimeout
	endPoint := fmt.Sprintf(":%s", config.ServerConfig.HttpPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
