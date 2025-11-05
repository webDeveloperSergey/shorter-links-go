package main

import (
	"fmt"
	"link-shorter/configs"
	"link-shorter/internal/auth"
	"link-shorter/internal/link"
	"link-shorter/pkg/db"
	"net/http"
)


func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{})

	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on 8081 port")

	server.ListenAndServe()
	
}




