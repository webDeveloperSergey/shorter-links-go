package auth

import (
	"fmt"
	"net/http"
)

type AuthHandler struct {}

func NewAuthHandler(router *http.ServeMux) {
	handler := &AuthHandler{}
	router.HandleFunc("POST /register", handler.Register())
	router.HandleFunc("POST /login", handler.Login())
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Login")
	}
}