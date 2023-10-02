package route

import "github.com/Anupam-dagar/baileys/server"

func SetupRoutes() {
	rootRouterGroup := server.GetGinEngine().GetRootRouterGroup()

	server.AddRoute(BPollRoutes(rootRouterGroup))
	server.AddRoute(BPollOptionRoutes(rootRouterGroup))
	server.AddRoute(BVoteRoutes(rootRouterGroup))
}
