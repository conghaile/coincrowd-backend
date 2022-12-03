package main

import (
	"log"

	"github.com/conghaile/coincrowd-API/api"
	"github.com/conghaile/coincrowd-API/db"
)

func main() {
	store, err := db.NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":3000", store)
	server.Run()
}
