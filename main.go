package main

import (
	"github.com/amitesh080295/GoMicroserviceProject/internal/database"
	"github.com/amitesh080295/GoMicroserviceProject/internal/server"
	"github.com/labstack/gommon/log"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("Failed to initialize database client: %s", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatalf(err.Error())
	}
}
