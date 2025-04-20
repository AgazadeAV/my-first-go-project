package bootstrap

import (
	"github.com/AgazadeAV/my-first-go-project/ent"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/handler"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/repository"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/service"
)

func InitUserModule(client *ent.Client) *handler.Handler {
	repo := repository.NewRepository(client)
	svc := service.NewService(repo)
	return handler.NewHandler(svc)
}
