package http

import (
	"github.com/gin-gonic/gin"
	"musicService/internal/usecases"
)

type UserHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHandler(r *gin.Engine, u usecases.UserUsecase) {
	handler := &UserHandler{
		userUsecase: u,
	}

	r.GET("/users", handler.GetAllUsers)
}
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(200, users)
}
