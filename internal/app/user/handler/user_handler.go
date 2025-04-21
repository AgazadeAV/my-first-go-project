package handler

import (
	"context"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/errs"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/model"
	"github.com/AgazadeAV/my-first-go-project/internal/app/user/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

const (
	UserBasePath = "/users"

	CreateUserRoute  = "/create-user"
	GetAllUsersRoute = "/get-all-users"
	DeleteUserRoute  = "/delete-user/:id"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (handler *Handler) RegisterRoutes(engine *gin.Engine) {
	users := engine.Group(UserBasePath)
	{
		users.POST(CreateUserRoute, handler.createUser)
		users.GET(GetAllUsersRoute, handler.getAllUsers)
		users.DELETE(DeleteUserRoute, handler.deleteUser)
	}
}

// CreateUser godoc
// @Summary Create new user
// @Description Add a new user to the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.CreateUserInput true "User data"
// @Success 200 {object} model.UserResponse
// @Failure 400 {object} errs.ErrorResponse
// @Failure 500 {object} errs.ErrorResponse
// @Router /users/create-user [post]
func (handler *Handler) createUser(ctx *gin.Context) {
	var input model.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		_ = ctx.Error(errs.ErrBadJSON)
		return
	}

	response, err := handler.service.CreateUser(context.Background(), input)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve all users from the database
// @Tags users
// @Produce json
// @Success 200 {array} model.UserResponse
// @Failure 500 {object} errs.ErrorResponse
// @Router /users/get-all-users [get]
func (handler *Handler) getAllUsers(ctx *gin.Context) {
	response, err := handler.service.GetAllUsers(context.Background())
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user by UUID
// @Tags users
// @Param id path string true "User UUID"
// @Produce json
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} errs.ErrorResponse
// @Failure 404 {object} errs.ErrorResponse
// @Router /users/delete-user/{id} [delete]
func (handler *Handler) deleteUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		_ = ctx.Error(errs.ErrBadUUID)
		return
	}

	if err := handler.service.DeleteUser(context.Background(), id); err != nil {
		_ = ctx.Error(err)
		return
	}

	response := model.SuccessResponse{Message: "User deleted"}

	ctx.JSON(http.StatusOK, response)
}
