package user

import (
	"context"

	"github.com/AgazadeAV/my-first-go-project/ent"
	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, input CreateUserInput) (*ent.User, error) {
	return s.repo.Create(ctx, input)
}

func (s *Service) GetAllUsers(ctx context.Context) ([]*ent.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteByID(ctx, id)
}
