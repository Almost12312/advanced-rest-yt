package main

import (
	"log"
	"net"
	"net/http"
	"time"
	
	"advanced-rest-yt/internal/user"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	log.Println("create router success!")

	userHandler := user.NewHandler()
	userHandler.Register(router)

	start(router)
}

func start(r *httprouter.Router) {

	l, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalln("Cant create logger", err)
	}

	log.Println("logger success created!")

	server := &http.Server{
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started!")
	err = server.Serve(l)
	if err != nil {
		log.Fatalln("Server was closed!", err)
	}

}
