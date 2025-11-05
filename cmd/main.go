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
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(db)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on 8081 port")

	server.ListenAndServe()
	
}




