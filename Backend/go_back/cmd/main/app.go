package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"training/pkg/handlers"
	"training/src/contodb"
)

const (
	createUser = "/users/create"
	allUsers   = "/users"
	userById   = "/users/:id"
	allTracks  = "/tracks"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}

	db, err := contodb.ConnectToDb()
	if err != nil {
		if err != nil {
			log.Fatalf("Ошибка при подключении к базе данных: %v", err)
		}
	}
	defer db.Close()

	r := gin.Default()

	r.POST(createUser, handlers.CreateUserHandler(db))

	r.Run(":5252")
}
