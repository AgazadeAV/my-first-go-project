package mapper

import (
	"github.com/AgazadeAV/my-first-go-project/ent"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/model"
)

func ToUserResponse(user *ent.User) *model.UserResponse {
	return &model.UserResponse{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		BirthDate:   user.BirthDate.Format("2006-01-02"),
	}
}

func ToUserResponseList(users []*ent.User) []*model.UserResponse {
	var result []*model.UserResponse
	for _, user := range users {
		result = append(result, ToUserResponse(user))
	}
	return result
}
