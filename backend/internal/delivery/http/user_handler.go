package http

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"musicService/internal/entities"
	"musicService/internal/usecases"
	"strconv"
)

type UserHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHandler(r *gin.Engine, u usecases.UserUsecase) {
	handler := &UserHandler{
		userUsecase: u,
	}

	r.GET("/users", handler.GetAllUsers)
	r.GET("/users/:id", handler.GetUserByID)
	r.POST("/users", handler.CreateUser)
}
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(200, users)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userUsecase.GetUserByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Error fetching user"})
		return
	}
	c.JSON(200, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var NewUser entities.User
	if err := c.ShouldBindJSON(&NewUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	userID, err := h.userUsecase.CreateUser(NewUser)
	if err != nil {
		c.JSON(500, gin.H{"error": "error creating user"})
		return
	}
	c.JSON(201, gin.H{"id": userID})
}
