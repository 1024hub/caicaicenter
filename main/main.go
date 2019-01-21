package main

import (
	"github.com/gin-gonic/gin"
 	"yibiancheng.com/accountcenter/routers"
	"yibiancheng.com/accountcenter/service"

)
func main() {
	service.Initdb()  //mysql 链接
	service.InitRedis() //redis
	gin.SetMode(gin.DebugMode)
	router := routers.InitRouter()
	router.Run()
}