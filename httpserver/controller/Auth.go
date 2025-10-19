package controller

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"mc-server-manager/conf"
	"net/http"
	"net/url"
	"strings"
)

type Claims struct {
	//Email string `json:"email"` // custom field
	jwt.RegisteredClaims
}

func AddAuthMiddleware(router *mux.Router) {
	middleware := mux.MiddlewareFunc(createAuthMiddlewareHandler)
	router.Use(middleware)
}

func createAuthMiddlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//// Skip auth for login and callback
		if r.URL.Path == "/login" || r.URL.Path == "/callback" {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims := Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.Conf.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Attach email to context for later use
		ctx := context.WithValue(r.Context(), "userEmail", "TODO_EMAIL")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func BindAuthEndpoints(router *mux.Router) {
	router.HandleFunc("/login", handleLogin).Methods("GET")
	router.HandleFunc("/callback", handleCallback).Methods("GET")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	authUrl := conf.GoogleOauth.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.SetAuthURLParam("prompt", "select_account"))
	http.Redirect(w, r, authUrl, http.StatusTemporaryRedirect)
}

// TODO redirect back to login with failed msg???
func handleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	token, err := conf.GoogleOauth.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Token exchange failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	idTokenRaw, _ := token.Extra("id_token").(string)
	claims := jwt.MapClaims{}
	_, _, err = new(jwt.Parser).ParseUnverified(idTokenRaw, claims)
	if err != nil {
		http.Error(w, "Token parse failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var user struct {
		Email     string `json:"email"`
		Name      string `json:"name"`
		FirstName string `json:"given_name"`
		LastName  string `json:"family_name"`
		Picture   string `json:"picture"`
	}

	user.Email = claims["email"].(string)
	user.Name = claims["name"].(string)
	user.FirstName = claims["given_name"].(string)
	user.LastName = claims["family_name"].(string)

	//TODO add custom claims!
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	tokenString, err4 := jwtToken.SignedString([]byte(conf.Conf.JwtSecret))
	if err4 != nil {
		http.Error(w, "Failed to sign JWT: "+err4.Error(), http.StatusInternalServerError)
		return
	}

	redirectURL := fmt.Sprintf("%s?token=%s", conf.Conf.AuthSuccessUrl, url.QueryEscape(tokenString))
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}
