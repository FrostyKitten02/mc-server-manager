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
	"time"
)

// TODO: add refresh token to jwt?
type Claims struct {
	User User `json:"user"`
	jwt.RegisteredClaims
}

type User struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func GetCallingUser(r *http.Request) User {
	return r.Context().Value("user").(User)
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

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.Conf.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims.User)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// TODO: update paths!!!
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

	currentTime := time.Now()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		User: User{
			Email:     claims["email"].(string),
			Name:      claims["name"].(string),
			FirstName: claims["given_name"].(string),
			LastName:  claims["family_name"].(string),
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "mc-server-manager",
			Subject:  claims["email"].(string),
			Audience: []string{"mc-server-manager-web"},
			ExpiresAt: &jwt.NumericDate{
				Time: currentTime.Add(24 * time.Hour),
			},
			IssuedAt: &jwt.NumericDate{
				Time: currentTime,
			},
			NotBefore: &jwt.NumericDate{
				Time: currentTime,
			},
		},
	})
	tokenString, err4 := jwtToken.SignedString([]byte(conf.Conf.JwtSecret))
	if err4 != nil {
		http.Error(w, "Failed to sign JWT: "+err4.Error(), http.StatusInternalServerError)
		return
	}

	redirectURL := fmt.Sprintf("%s?token=%s", conf.Conf.AuthSuccessUrl, url.QueryEscape(tokenString))
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}
