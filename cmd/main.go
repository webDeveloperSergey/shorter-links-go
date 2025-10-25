package main

import (
	"link-shorter/configs"
	"fmt"
	"link-shorter/internal/auth"
	"net/http"
)


func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on 8081 port")

	server.ListenAndServe()
	
}




