package main

import (
	"github.com/JonasBrothers12/BackendChallenge/routes"
	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.POST("/CreateAccount",routes.HandleCreateAccount)
  r.Run() 
}