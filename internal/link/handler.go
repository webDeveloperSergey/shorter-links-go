package link

import (
	"link-shorter/configs"
	"net/http"
)

type LinkHandlerDeps struct {

}

type LinkHandler struct {
	*configs.Config
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{}

	router.HandleFunc("GET {hash}", handler.GoTo())
	router.HandleFunc("POST /link/", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}
