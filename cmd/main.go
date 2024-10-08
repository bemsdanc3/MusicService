package main

import (
	"github.com/gin-gonic/gin"
	"musicService/internal/delivery/http"
	"musicService/internal/repository"
	"musicService/internal/usecases"
	"musicService/pkg/db"
)

func main() {
	r := gin.Default()

	// Подключение к базе данных
	dbConn, err := db.Connect()
	if err != nil {
		panic("Failed to connect to database")
	}

	// Создаем репозиторий и use case
	userRepo := repository.NewUserRepository(dbConn)
	userUsecase := usecases.NewUserUsecase(userRepo)

	// Создаем обработчик для пользователей
	http.NewUserHandler(r, userUsecase)

	r.Run(":5002")
}
