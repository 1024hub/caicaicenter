package test2

import (
	"github.com/gin-gonic/gin"
	"yibiancheng.com/accountcenter/controller"
)

func Routes(route *gin.Engine){
	test2:=route.Group("/test2")
	test2.GET("/get", controller.Test2GetApi)
	test2.POST("/post", controller.Test2PostApi)
}