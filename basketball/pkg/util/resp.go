package util

import (
	"github.com/gin-gonic/gin"
)

func Output(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(200, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
