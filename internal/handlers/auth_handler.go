package handlers

import "github.com/gin-gonic/gin"

type authHandler struct {
}

func NewAuthHandler() authHandler {
	return authHandler{}
}

func (auth *authHandler) SampleRouter(c *gin.Context) {

	c.JSON(200, "message: success")
}
