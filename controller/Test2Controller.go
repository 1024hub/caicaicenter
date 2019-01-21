package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test2GetApi(c *gin.Context) {
	c.JSON(http.StatusOK,"hello test2_controller_getapi")
}
func Test2PostApi(c *gin.Context) {
	c.JSON(http.StatusOK,"hello test2_controller_postapi")
}
