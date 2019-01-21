package test1

import (
	"github.com/gin-gonic/gin"
	"yibiancheng.com/accountcenter/controller"
)

func Routes(route *gin.Engine){
	test1:=route.Group("/test1")
 	test1.GET("/get", controller.Test1GetApi)
	test1.POST("/post", controller.Test1PostApi)
}