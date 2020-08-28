package v1

import (
	"github.com/gin-gonic/gin"
	"goworker/core/app"
	"goworker/core/e"
	"net/http"
)

func SayHello(ctx *gin.Context)  {
	gapp := app.Gin{G: ctx}
	gapp.Response(http.StatusOK, e.SUCCESS, "hello, gowrker")
}