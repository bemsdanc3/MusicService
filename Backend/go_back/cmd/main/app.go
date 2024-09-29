package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	db2 "training/pkg/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}

	db, err := db2.ConnectToDb()
	if err != nil {
		if err != nil {
			log.Fatalf("Ошибка при подключении к базе данных: %v", err)
		}
	}
	defer db.Close()

	r := gin.Default()

	r.Run(":1337")
}
