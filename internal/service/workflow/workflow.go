package workflow

import (
	"github.com/gin-gonic/gin"
	"workflow_http/internal/service"
)

func (s *Service) Test(c *gin.Context) {
	service.GetResponse(c, 0, "success", nil)
}
