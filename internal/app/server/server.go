package server

import (
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/errs"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/handler"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(userHandler *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.Use(errs.ErrorHandler())

	userHandler.RegisterRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
