package workflow

import "github.com/gin-gonic/gin"

type IWorkflow interface {
	Test(c *gin.Context)
}
