package route

import (
	"github.com/Anupam-dagar/baileys/constant/types"
	"github.com/gin-gonic/gin"
	"polls/controller"
)

func PollOptionRoutes(routerGroup *gin.RouterGroup) {
	router := routerGroup.Group("/poll-options")
	{
		controller.NewPollOptionController(router)
	}
}

func BPollOptionRoutes(routerGroup *gin.RouterGroup) types.RouteFunc {
	return func() (*gin.RouterGroup, func(rg *gin.RouterGroup)) {
		return routerGroup, PollOptionRoutes
	}
}
