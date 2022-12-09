package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}
