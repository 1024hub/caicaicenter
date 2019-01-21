package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test1GetApi(c *gin.Context) {
	c.JSON(http.StatusOK,"hello test1_controller_getapi")
}
func Test1PostApi(c *gin.Context) {
	c.JSON(http.StatusOK,"hello test1_controller_postapi")
}
