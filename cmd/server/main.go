package main

import (
	"log"
	"go-threads/internal/app/server"
)

func main() {
	server := server.NewServer(100)
	if err := server.Start(":12345"); err != nil {
		log.Fatal(err)
	}
}
