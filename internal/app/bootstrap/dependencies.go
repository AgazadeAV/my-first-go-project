package bootstrap

import (
	"github.com/AgazadeAV/my-first-go-project/ent"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/handler"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/repository"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/service"
	"github.com/AgazadeAV/my-first-go-project/internal/workerpool"
)

func InitUserModule(client *ent.Client, pool *workerpool.Pool) *handler.Handler {
	repo := repository.NewRepository(client)
	svc := service.NewService(repo, pool)
	return handler.NewHandler(svc)
}
