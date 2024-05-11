package middleware

import (
	"github.com/JonasBrothers12/BackendChallenge/config"
	"github.com/gin-gonic/gin"
)

func SetMiddlewares(r *gin.Engine,cfg *config.Config) {
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
}