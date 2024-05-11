package server

import (
	"github.com/JonasBrothers12/BackendChallenge/config"

	middleware "github.com/JonasBrothers12/BackendChallenge/server/Middleware"
	"github.com/JonasBrothers12/BackendChallenge/server/routes"
	"github.com/gin-gonic/gin"
)

func InitializingServer(cfg *config.Config) {
  r := gin.New()  
  middleware.SetMiddlewares(r,cfg)
  router.Routes(r)
  r.Run(":8080") 
}
