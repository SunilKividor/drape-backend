package handlers

import (
	"log"
	"net/http"

	"github.com/SunilKividor/drape/internal/auth"
	"github.com/gin-gonic/gin"
)

type authHandler struct{}

func NewAuthHandler() authHandler {
	return authHandler{}
}

func (authentication *authHandler) Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "successful login",
	})
}

func (authentication *authHandler) OAuthGoogleLogin(c *gin.Context) {
	authCode := auth.OAuthGoogleLogin()
	c.Redirect(http.StatusTemporaryRedirect, authCode)
}

func (authentication *authHandler) OAuthGoogleCallback(c *gin.Context) {
	code := c.Query("code")
	userInfo := auth.OAuthGoogleCallback(code)
	log.Println(userInfo)
	c.Redirect(http.StatusTemporaryRedirect, "/home")
}
