package builder

import (
	"github.com/AgazadeAV/my-first-go-project/ent"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/model"
	"time"
)

type Builder struct {
	input model.CreateUserInput
}

func NewUserBuilder(input model.CreateUserInput) *Builder {
	return &Builder{input: input}
}

func (userBuilder *Builder) Build(create *ent.UserCreate) (*ent.UserCreate, error) {
	birthDate, _ := time.Parse("2006-01-02", *userBuilder.input.BirthDate)

	return create.
		SetFirstName(*userBuilder.input.FirstName).
		SetLastName(*userBuilder.input.LastName).
		SetUsername(*userBuilder.input.Username).
		SetEmail(*userBuilder.input.Email).
		SetPhoneNumber(*userBuilder.input.PhoneNumber).
		SetBirthDate(birthDate), nil
}
