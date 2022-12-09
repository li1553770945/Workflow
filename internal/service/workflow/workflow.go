package workflow

import (
	"HomeWorkGo/internal/service"
	"github.com/gin-gonic/gin"
)

func (s *Service) Test(c *gin.Context) {
	service.GetResponse(c, 0, "success", nil)
}
