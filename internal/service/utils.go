package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"workflow_http/internal/constant"
)

func GetResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func SplitMessage(message string) (name string, content string, err error) {
	err = nil
	name, content, ok := strings.Cut(message, " ")
	if !ok {
		err = errors.New(constant.MessageFormatError)
	}
	return
}
