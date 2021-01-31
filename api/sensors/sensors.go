package sensors

import (
	"github.com/gin-gonic/gin"
	"pi-hub/api/sensors/security"
)

func RoutesHandler(g *gin.RouterGroup) {
	security.RoutesHandler(g.Group("/security"))
}
