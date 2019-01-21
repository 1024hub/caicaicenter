package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
 	"net/http"
 	"yibiancheng.com/accountcenter/entity"
 	"yibiancheng.com/accountcenter/service"
)




func IndexApi(c *gin.Context) {
	 c.JSON(http.StatusOK,"hello gin")
}


// 添加一条user表数据
func Insert(c *gin.Context)  {
	fmt.Println("进入controller")
	user:=&entity.User{
		Username:"chenbo",
	}
	service.SaveUsers(user)
}
func Update(c *gin.Context) {
	c.JSON(http.StatusOK,"hello gin")
}
func Del(c *gin.Context) {
	c.JSON(http.StatusOK,"hello gin")
}
func Find(c *gin.Context) {
	c.JSON(http.StatusOK,"hello gin")
}
