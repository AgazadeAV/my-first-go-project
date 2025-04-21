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

func (b *Builder) Build(create *ent.UserCreate) (*ent.UserCreate, error) {
	birthDate, _ := time.Parse("2006-01-02", *b.input.BirthDate)

	return create.
		SetFirstName(*b.input.FirstName).
		SetLastName(*b.input.LastName).
		SetUsername(*b.input.Username).
		SetEmail(*b.input.Email).
		SetPhoneNumber(*b.input.PhoneNumber).
		SetBirthDate(birthDate), nil
}
