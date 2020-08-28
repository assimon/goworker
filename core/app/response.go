package app

import (
	"github.com/gin-gonic/gin"
	"goworker/core/e"
)

type Gin struct {
	G *gin.Context
}

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(httpCode, errCode int, data interface{})  {
	g.G.JSON(httpCode, Response{
		Code: httpCode,
		Msg: e.GetMessage(errCode),
		Data: data,
	})
	return
}