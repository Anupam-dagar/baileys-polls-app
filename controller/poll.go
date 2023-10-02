package controller

import (
	"github.com/Anupam-dagar/baileys/controller"
	"github.com/gin-gonic/gin"
	"polls/entity"
)

type PollControllerInterface interface {
	controller.BaseControllerInterface
}

type pollController struct {
	controller.BaseControllerInterface
}

func NewPollController(rg *gin.RouterGroup) PollControllerInterface {
	pc := new(pollController)
	pc.BaseControllerInterface = controller.NewBaseController[entity.PollPtr](rg)

	return pc
}
