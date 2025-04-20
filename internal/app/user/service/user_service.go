package service

import (
	"context"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/builder"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/model"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/repository"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/service/mapper"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/util"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/validation"
	"github.com/google/uuid"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (service *Service) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.UserResponse, error) {
	if err := validation.ValidateUserInputFields(input); err != nil {
		return nil, err
	}

	if err := service.ValidateUniqueness(ctx, input); err != nil {
		return nil, err
	}

	userBuilder := builder.NewUserBuilder(input)
	create, err := userBuilder.Build(service.repo.Client().User.Create())
	if err != nil {
		return nil, err
	}

	user, err := service.repo.Create(ctx, create)
	if err != nil {
		return nil, err
	}

	return mapper.ToUserResponse(user), nil
}

func (service *Service) GetAllUsers(ctx context.Context) ([]*model.UserResponse, error) {
	users, err := service.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.ToUserResponseList(users), nil
}

func (service *Service) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return service.repo.DeleteByID(ctx, id)
}

func (service *Service) ValidateUniqueness(ctx context.Context, input model.CreateUserInput) error {
	fieldErrors := make(validation.FieldErrors)

	if exists, _ := service.repo.IsEmailTaken(ctx, *input.Email); exists {
		fieldErrors["email"] = util.ErrEmailAlreadyExists
	}
	if exists, _ := service.repo.IsUsernameTaken(ctx, *input.Username); exists {
		fieldErrors["username"] = util.ErrUsernameAlreadyExists
	}
	if exists, _ := service.repo.IsPhoneTaken(ctx, *input.PhoneNumber); exists {
		fieldErrors["phone_number"] = util.ErrPhoneAlreadyExists
	}

	if len(fieldErrors) > 0 {
		return fieldErrors
	}
	return nil
}
