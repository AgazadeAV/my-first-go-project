package validation

import (
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/model"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/util"
	"regexp"
	"strings"
	"time"
)

type FieldErrors map[string]string

func (fe FieldErrors) Error() string {
	return "Validation error"
}

var (
	nameRegex     = regexp.MustCompile(util.NamePattern)
	usernameRegex = regexp.MustCompile(util.UsernamePattern)
	emailRegex    = regexp.MustCompile(util.EmailPattern)
	phoneRegex    = regexp.MustCompile(util.PhonePattern)
)

func ValidateUserInputFields(input model.CreateUserInput) error {
	fieldErrors := make(FieldErrors)

	// First Name
	if msg, ok := validateStringField(input.FirstName, util.ErrFirstNameRequired, util.ErrFirstNameEmpty, util.ErrFirstNameInvalid, nameRegex, util.NameMinLength, util.NameMaxLength); !ok {
		fieldErrors["first_name"] = msg
	}

	// Last Name
	if msg, ok := validateStringField(input.LastName, util.ErrLastNameRequired, util.ErrLastNameEmpty, util.ErrLastNameInvalid, nameRegex, util.NameMinLength, util.NameMaxLength); !ok {
		fieldErrors["last_name"] = msg
	}

	// Username
	if msg, ok := validateStringField(input.Username, util.ErrUsernameRequired, util.ErrUsernameEmpty, util.ErrUsernameInvalid, usernameRegex, util.UserNameMinLength, util.UserNameMaxLength); !ok {
		fieldErrors["username"] = msg
	}

	// Email
	if msg, ok := validateStringField(input.Email, util.ErrEmailRequired, util.ErrEmailEmpty, util.ErrEmailInvalid, emailRegex, util.EmailMinLength, util.EmailMaxLength); !ok {
		fieldErrors["email"] = msg
	}

	// Phone
	if msg, ok := validateStringField(input.PhoneNumber, util.ErrPhoneRequired, util.ErrPhoneEmpty, util.ErrPhoneInvalid, phoneRegex, util.PhoneNumberLength, util.PhoneNumberLength); !ok {
		fieldErrors["phone_number"] = msg
	}

	// Birth Date
	if input.BirthDate == nil {
		fieldErrors["birth_date"] = util.ErrBirthDateRequired
	} else if strings.TrimSpace(*input.BirthDate) == "" {
		fieldErrors["birth_date"] = util.ErrBirthDateEmpty
	} else {
		parsed, err := time.Parse("2006-01-02", *input.BirthDate)
		if err != nil {
			fieldErrors["birth_date"] = util.ErrBirthDateInvalid
		} else if reason := validateBirthDate(parsed); reason != "" {
			fieldErrors["birth_date"] = reason
		}
	}

	if len(fieldErrors) > 0 {
		return fieldErrors
	}
	return nil
}

func validateStringField(
	field *string,
	requiredMsg, emptyMsg, formatMsg string,
	pattern *regexp.Regexp,
	min, max int,
) (string, bool) {
	if field == nil {
		return requiredMsg, false
	}
	value := strings.TrimSpace(*field)
	if value == "" {
		return emptyMsg, false
	}
	if len(value) < min || len(value) > max || !pattern.MatchString(value) {
		return formatMsg, false
	}
	return "", true
}

func validateBirthDate(birthDate time.Time) string {
	minDate := time.Date(util.MinimalBirthYear, 1, 1, 0, 0, 0, 0, time.UTC)
	today := time.Now()
	adultAgeDate := today.AddDate(-util.MinimalUserAge, 0, 0)

	if birthDate.After(today) {
		return util.ErrBirthDateInFuture
	}
	if birthDate.Before(minDate) {
		return util.ErrBirthDateTooEarly
	}
	if birthDate.After(adultAgeDate) {
		return util.ErrBirthDateTooYoung
	}
	return ""
}
