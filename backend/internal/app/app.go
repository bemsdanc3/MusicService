package app

import (
	"github.com/gin-gonic/gin"
	"musicService/internal/delivery/http"
	"musicService/internal/repository"
	"musicService/internal/usecases"
	"musicService/pkg/db"
)

type App struct {
	Router *gin.Engine
}

func New() *App {
	app := &App{
		Router: gin.Default(),
	}

	// Подключение к базе данных
	dbConn, err := db.Connect()
	if err != nil {
		panic("Failed to connect to database")
	}

	// Создаем репозиторий и use case
	userRepo := repository.NewUserRepository(dbConn)
	userUsecase := usecases.NewUserUsecase(userRepo)

	// Создаем обработчики
	http.NewUserHandler(app.Router, userUsecase)

	return app
}
