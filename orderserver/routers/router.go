package routers

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func InitRouters() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.GET("/orders/", func(context *gin.Context) {
		//context.String(200, "get orderinfos")
		context.JSON(200, Response{
			Code: 200,
			Msg:  "get order success",
			Data: nil,
		})
		return
	})
	return ginRouter
}