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
	allUsers          = "/users"
	userById          = "/users/:id"
	allArtists        = "/artists"
	artistById        = "/artists/:id"
	allAlbums         = "albums"
	albumsById        = "/albums/:id"
	allGenres         = "/genres"
	genresById        = "/genres/:id"
	allTracks         = "/tracks"
	tracksById        = "/tracks/:id"
	allPlaylists      = "/playlists"
	playlistById      = "/playlists/:id"
	addTrackById      = "/playlists/:id/tracks"
	deleteTrackById   = "/playlists/:id/tracks/:id"
	allListeningStory = "/listening_history"
	getListeningStory = "/listening_history/user/:id"
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

	r.POST(allUsers, handlers.CreateUserHandler(db))

	r.GET(allUsers, handlers.GetAllUsersHandler(db))
	r.GET(userById, handlers.GetUserByIdHandler(db))

	r.PUT(userById, handlers.UpdateUserHandler(db))

	r.Run(":5252")
}
