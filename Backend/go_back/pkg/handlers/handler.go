package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
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
