package security

import "github.com/gin-gonic/gin"

func RoutesHandler(g *gin.RouterGroup) {
	DoorsRoutesHandler(g.Group("/doors"))
}
