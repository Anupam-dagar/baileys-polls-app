package controller

import (
	"github.com/Anupam-dagar/baileys/controller"
	"github.com/gin-gonic/gin"
	"polls/entity"
)

type VoteControllerInterface interface {
	controller.BaseControllerInterface
}

type voteController struct {
	controller.BaseControllerInterface
}

func NewVoteController(rg *gin.RouterGroup) VoteControllerInterface {
	vc := new(voteController)
	vc.BaseControllerInterface = controller.NewBaseController[entity.VotePtr](rg)

	return vc
}
