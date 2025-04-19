package main

// @title Users API
// @version 1.0
// @description This is a simple REST API with Gin and Ent
// @host localhost:8080
// @BasePath /

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/AgazadeAV/my-first-go-project/ent"
	"github.com/AgazadeAV/my-first-go-project/internal/user"

	_ "github.com/AgazadeAV/my-first-go-project/docs"
	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=user dbname=usersdb password=password sslmode=disable search_path=myschema")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	repo := user.NewRepository(client)
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	router := gin.Default()
	handler.RegisterRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("🚀 Server started at http://localhost:8080")
	router.Run(":8080")
}
