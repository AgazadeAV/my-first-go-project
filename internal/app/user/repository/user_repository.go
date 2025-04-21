package repository

import (
	"context"
	"github.com/AgazadeAV/my-first-go-project/ent"
	"github.com/AgazadeAV/my-first-go-project/ent/user"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/errs"
	"github.com/google/uuid"
)

type Repository struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) *Repository {
	return &Repository{client: client}
}

func (r *Repository) Client() *ent.Client {
	return r.client
}

func (r *Repository) Create(ctx context.Context, userCreate *ent.UserCreate) (*ent.User, error) {
	return userCreate.Save(ctx)
}

func (r *Repository) GetAll(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

func (r *Repository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	err := r.client.User.DeleteOneID(id).Exec(ctx)
	if ent.IsNotFound(err) {
		return errs.ErrNotFound
	}
	return err
}

func (r *Repository) IsEmailTaken(ctx context.Context, email string) (bool, error) {
	return r.client.User.Query().Where(user.EmailEQ(email)).Exist(ctx)
}

func (r *Repository) IsPhoneTaken(ctx context.Context, phone string) (bool, error) {
	return r.client.User.Query().Where(user.PhoneNumberEQ(phone)).Exist(ctx)
}

func (r *Repository) IsUsernameTaken(ctx context.Context, username string) (bool, error) {
	return r.client.User.Query().Where(user.UsernameEQ(username)).Exist(ctx)
}
