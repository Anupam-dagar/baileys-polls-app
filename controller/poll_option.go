package controller

import (
	"github.com/Anupam-dagar/baileys/controller"
	"github.com/gin-gonic/gin"
	"polls/entity"
)

type PollOptionControllerInterface interface {
	controller.BaseControllerInterface
}

type pollOptionController struct {
	controller.BaseControllerInterface
}

func NewPollOptionController(rg *gin.RouterGroup) PollOptionControllerInterface {
	poc := new(pollOptionController)
	poc.BaseControllerInterface = controller.NewBaseController[entity.PollOptionPtr](rg)

	return poc
}
