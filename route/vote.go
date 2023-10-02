package route

import (
	"github.com/Anupam-dagar/baileys/constant/types"
	"github.com/gin-gonic/gin"
	"polls/controller"
)

func VoteRoutes(routerGroup *gin.RouterGroup) {
	router := routerGroup.Group("/votes")
	{
		controller.NewVoteController(router)
	}
}

func BVoteRoutes(routerGroup *gin.RouterGroup) types.RouteFunc {
	return func() (*gin.RouterGroup, func(rg *gin.RouterGroup)) {
		return routerGroup, VoteRoutes
	}
}
