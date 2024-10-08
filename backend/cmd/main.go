package main

import "musicService/internal/app"

func main() {
	a := app.New()
	a.Router.Run(":5002")
}
