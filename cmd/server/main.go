package main

import (
	"github.com/AgazadeAV/my-first-go-project/internal/app/bootstrap"
	"github.com/AgazadeAV/my-first-go-project/internal/app/server"
	"github.com/AgazadeAV/my-first-go-project/pkg/database"
	"log"
)

func main() {
	client := database.NewEntClient()
	defer client.Close()

	userHandler := bootstrap.InitUserModule(client)
	router := server.NewRouter(userHandler)

	log.Println("🚀 Server started at http://localhost:8080")
	router.Run(":8080")
}
