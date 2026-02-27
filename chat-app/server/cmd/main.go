package main

import (
	"log"

	"github.com/Harman6282/chat-app/db"
	"github.com/Harman6282/chat-app/internal/user"
	"github.com/Harman6282/chat-app/router"
)


func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error connecting db: %v", err)
	}


	userRepo :=	user.NewRepository(db.GetDB())
	userSvc := user.NewService(userRepo)
	userHandler := user.NewHandler(userSvc)


	router.InitRouter(userHandler)
	router.Start(":8080")

}