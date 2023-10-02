package route

import (
	"github.com/Anupam-dagar/baileys/constant/types"
	"github.com/gin-gonic/gin"
	"polls/controller"
)

func PollRoutes(routerGroup *gin.RouterGroup) {
	router := routerGroup.Group("/polls")
	{
		controller.NewPollController(router)
	}
}

func BPollRoutes(routerGroup *gin.RouterGroup) types.RouteFunc {
	return func() (*gin.RouterGroup, func(rg *gin.RouterGroup)) {
		return routerGroup, PollRoutes
	}
}
