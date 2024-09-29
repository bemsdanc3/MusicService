package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"
	"training/pkg/requests"
)

func CreateUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user requests.User

		//вытаскиваю данные из тела запроса
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		//проверка на наличие логина, почты и пароля
		if user.Login == "" || user.Email == "" || user.Pass == "" {
			c.JSON(400, gin.H{"error": "Missing login, email or password"})
			return
		}

		//создание самого пользователя
		err := requests.CreateUser(db, user)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"message": "User created successfully!"})
	}

}

func GetAllUsersHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := requests.GetAllUsers(db)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, users)
	}

}

func GetUserByIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
		}

		user, err := requests.GetUserById(db, uint(id))
		if err != nil {
			if err.Error() == "user not found" {
				c.JSON(404, gin.H{"error": err.Error()})
			} else {
				c.JSON(500, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(200, user)
	}
}

func UpdateUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user requests.User

		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}

		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		err = requests.UpdateUser(db, uint(id), user)
		if err != nil {
			if err.Error() == "user not found" {
				c.JSON(404, gin.H{"error": err.Error()})
			} else {
				c.JSON(500, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(200, gin.H{"message": "User updated successfully!"})
	}
}
