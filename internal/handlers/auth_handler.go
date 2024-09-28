package handlers

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/auth/credentials/idtoken"
	"github.com/SunilKividor/drape/internal/configs"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type authHandler struct{}

func NewAuthHandler() authHandler {
	return authHandler{}
}

func (auth *authHandler) Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "successful login",
	})
}

func (auth *authHandler) OAuthGoogleLogin(c *gin.Context) {
	oauthState := "sunilKividorRandomString@1296"
	u := configs.GoogleOauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	c.Redirect(http.StatusTemporaryRedirect, u)
}

func (auth *authHandler) OAuthGoogleCallback(c *gin.Context) {
	code := c.Query("code")
	userInfo := getCodeExchangeData(code)
	log.Println(userInfo)
	c.Redirect(http.StatusTemporaryRedirect, "/auth/success")
}

func getCodeExchangeData(code string) map[string]interface{} {
	token, err := configs.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Fatalf("error getting token")
	}
	idToken := token.Extra("id_token").(string)
	payload, err := idtoken.Validate(context.Background(), idToken, "")
	if err != nil {
		log.Fatalf("error validating idToken")
	}

	return map[string]interface{}{
		"userId": payload.Claims["sub"],
		"email":  payload.Claims["email"],
		"name":   payload.Claims["name"],
	}
}
