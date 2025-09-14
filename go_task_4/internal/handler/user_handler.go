package handler

import (
	"go_task_4/internal/model"
	"go_task_4/internal/response"
	"go_task_4/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"omitempty,email"`
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{userService: service}
}

func (u *UserHandler) Register(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request parameters")
		return
	}
	user := model.User{}
	user.Email = req.Email
	user.Username = req.Username
	user.Password = req.Password
	if err := u.userService.Register(&user); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (u *UserHandler) Login(c *gin.Context) {

	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request parameters")
		return
	}
	user := model.User{}
	user.Email = req.Email
	user.Username = req.Username
	user.Password = req.Password
	token, err := u.userService.Login(&user)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, token)
}
