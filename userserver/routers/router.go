package routers

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func InitRouters() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.GET("/users/", func(context *gin.Context) {
		//context.String(200, "get userinfos")
		context.JSON(200, Response{
			Code: 200,
			Msg:  "success",
			Data: nil,
		})
		return
	})
	//http://localhost:18001/users    {"code":200,"msg":"success","data":null}

	return ginRouter
}