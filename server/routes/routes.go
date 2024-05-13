package router

import (
	"github.com/JonasBrothers12/BackendChallenge/server/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(svr *gin.Engine,ctrl *controller.ControllerManager){
	svr.POST("/CreateAccount",ctrl.UserController.HandleNewUser)
}

