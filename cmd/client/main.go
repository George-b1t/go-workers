package main

import (
	"go-threads/internal/app/client"
	"log"
)

func main() {
	client := client.NewClient()

	if err := client.Connect("localhost:12345"); err != nil {
		log.Fatal(err)
	}

	client.Start()
}
