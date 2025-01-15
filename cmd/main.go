package main

import (
	"log"

	"github.com/Sahal-P/Go-Auth/cmd/api"
	"github.com/Sahal-P/Go-Auth/db"
)

func main() {
	dbStorage := db.NewPostgreSQLStorage()

	if err := dbStorage.Ping(); err != nil {
		log.Fatalf("Error pinging the DB: %v", err)
	}
	
	server := api.NewAPIServer(":8080", dbStorage)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
	
	defer dbStorage.Close()
}
