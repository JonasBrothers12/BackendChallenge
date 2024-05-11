package router

import (
	"github.com/JonasBrothers12/BackendChallenge/server/controller/user"
	"github.com/gin-gonic/gin"
)

func Routes(svr *gin.Engine){
	svr.POST("/CreateAccount",user.HandleNewuser)
}

