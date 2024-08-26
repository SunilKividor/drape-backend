package router

import (
	"github.com/SunilKividor/drape/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	authHandler := handlers.NewAuthHandler()
	router.GET("/sample", authHandler.SampleRouter)
}
