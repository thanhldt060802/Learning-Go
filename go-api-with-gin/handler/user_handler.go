package handler

import (
	"go-api-with-gin/dto"
	"go-api-with-gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (userHandler *UserHandler) GetAllUsers(context *gin.Context) {
	users, err := userHandler.userService.FindAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, users)
}

func (userHandler *UserHandler) GetUserById(context *gin.Context) {
	userId, _ := strconv.Atoi(context.Param("id"))

	user, err := userHandler.userService.FindUserById(int64(userId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (userHandler *UserHandler) GetUserByUsername(context *gin.Context) {
	username := context.Param("username")

	user, err := userHandler.userService.FindUserByUsername(username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (userHandler *UserHandler) CreateUser(context *gin.Context) {
	var newUser dto.CreateUserDTO

	if err := context.ShouldBind(&newUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := userHandler.userService.CreateUser(newUser); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, "Create user success!")
}

func (userHandler *UserHandler) UpdateUser(context *gin.Context) {
	userId, _ := strconv.Atoi(context.Param("id"))

	var updatingUser dto.UpdateUserDTO

	if err := context.ShouldBind(&updatingUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := userHandler.userService.UpdateUser(int64(userId), updatingUser); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, "Update user success!")
}

func (userHandler *UserHandler) DeleteUserById(context *gin.Context) {
	userId, _ := strconv.Atoi(context.Param("id"))

	if err := userHandler.userService.DeleteUserById(int64(userId)); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, "Delete user success!")
}
