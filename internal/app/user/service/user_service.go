package service

import (
	"context"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/builder"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/job"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/model"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/repository"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/service/mapper"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/util"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/validation"
	"github.com/AgazadeAV/my-first-go-project/internal/workerpool"
	"github.com/google/uuid"
)

type Service struct {
	repo *repository.Repository
	pool *workerpool.Pool
}

func NewService(repo *repository.Repository, pool *workerpool.Pool) *Service {
	return &Service{
		repo: repo,
		pool: pool,
	}
}

func (s *Service) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.UserResponse, error) {
	if err := validation.ValidateUserInputFields(input); err != nil {
		return nil, err
	}

	if err := s.ValidateUniqueness(ctx, input); err != nil {
		return nil, err
	}

	userBuilder := builder.NewUserBuilder(input)
	create, err := userBuilder.Build(s.repo.Client().User.Create())
	if err != nil {
		return nil, err
	}

	user, err := s.repo.Create(ctx, create)
	if err != nil {
		return nil, err
	}

	s.pool.Submit(job.WelcomeEmailJob{
		UserID: user.ID,
		Email:  user.Email,
	})

	return mapper.ToUserResponse(user), nil
}

func (s *Service) GetAllUsers(ctx context.Context) ([]*model.UserResponse, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.ToUserResponseList(users), nil
}

func (s *Service) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteByID(ctx, id)
}

func (s *Service) ValidateUniqueness(ctx context.Context, input model.CreateUserInput) error {
	fieldErrors := make(validation.FieldErrors)

	if exists, _ := s.repo.IsEmailTaken(ctx, *input.Email); exists {
		fieldErrors["email"] = util.ErrEmailAlreadyExists
	}
	if exists, _ := s.repo.IsUsernameTaken(ctx, *input.Username); exists {
		fieldErrors["username"] = util.ErrUsernameAlreadyExists
	}
	if exists, _ := s.repo.IsPhoneTaken(ctx, *input.PhoneNumber); exists {
		fieldErrors["phone_number"] = util.ErrPhoneAlreadyExists
	}

	if len(fieldErrors) > 0 {
		return fieldErrors
	}
	return nil
}
