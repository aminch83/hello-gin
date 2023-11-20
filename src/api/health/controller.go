package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (controller *HealthController) healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "Hello World\n")
}
