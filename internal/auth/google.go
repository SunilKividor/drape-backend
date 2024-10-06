package auth

import (
	"context"
	"log"

	"github.com/SunilKividor/drape/internal/configs"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
)

func OAuthGoogleLogin() string {
	oauthState := "sunilKividorRandomString@1296"
	cfg := configs.NewGoogleConfig()
	u := cfg.GoogleOauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	return u
}

func OAuthGoogleCallback(authCode string) interface{} {
	userInfo := getCodeExchangeData(authCode)
	log.Println(userInfo)
	return userInfo
}

func getCodeExchangeData(code string) map[string]interface{} {
	cfg := configs.NewGoogleConfig()
	token, err := cfg.GoogleOauthConfig.Exchange(context.Background(), code)
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
