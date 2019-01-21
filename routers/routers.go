package routers

import (
	"github.com/gin-gonic/gin"
	"yibiancheng.com/accountcenter/routers/test1"
	"yibiancheng.com/accountcenter/routers/test2"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	test1.Routes(router)  //对应test1control的路由都在这里
	test2.Routes(router)  //对应test2control的路由都在这里
	//taR := router.Group("/data")
	//taR.Use(middleware.Jwt())
	//
	//{
	//	taR.GET("/dataByTime", util.GetDataByTime)
	//}
	return router
}
