package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdin, "MYSERVER ", log.Lshortfile)
	myServer := server.NewServer(logger)
	
	if err := myServer.HttpServer.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
