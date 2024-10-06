package router

import (
	"github.com/SunilKividor/drape/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	authHandler := handlers.NewAuthHandler()
	//auth routes
	auth := router.Group("/auth")
	auth.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Success",
		})
	})
	auth.GET("/google/login", authHandler.OAuthGoogleLogin)
	auth.GET("/google/callback", authHandler.OAuthGoogleCallback)

	router.GET("/home", authHandler.Success)
}
