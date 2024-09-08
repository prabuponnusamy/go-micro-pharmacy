package main

import (
	"log"

	"mylearning.com/golang/micro/ln/pharmacy/service/database"
	"mylearning.com/golang/micro/ln/pharmacy/service/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("Unable to get database client: %s", err)
	}

	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatalf("Unable to start the server; %s", err)
	}
	log.Print("Server started successfully")
}
