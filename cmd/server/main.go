package main

// @title Users API
// @version 1.0
// @description This is a simple REST API with Gin and Ent
// @host localhost:8080
// @BasePath /

import (
	_ "github.com/AgazadeAV/my-first-go-project/docs"
	"github.com/AgazadeAV/my-first-go-project/internal/app/bootstrap"
	"github.com/AgazadeAV/my-first-go-project/internal/app/server"
	"github.com/AgazadeAV/my-first-go-project/pkg/database"
	"log"
)

func main() {
	client := database.NewEntClient()
	defer func() {
		if err := client.Close(); err != nil {
			log.Printf("Failed to close database client: %v", err)
		}
	}()

	userHandler := bootstrap.InitUserModule(client)
	router := server.NewRouter(userHandler)

	log.Println("🚀 Server started at http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
