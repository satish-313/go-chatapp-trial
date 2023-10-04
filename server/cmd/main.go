package main

import (
	"log"
	"server/db"
	"server/internal/user"
	"server/router"
)

func main() {
	dbConn, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("could not initialize db : %s ", err)
	}

	userRepo := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRepo)
	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start()
}
