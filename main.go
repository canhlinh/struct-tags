package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/zenazn/goji"
)

func main() {
	goji.Get("/", GetIndex)
	goji.Get("/user/:id", GetUser)
	goji.Post("/user", CreateUser)
	goji.Put("/user", UpdateUser)
	goji.Get("/config", GetConfig)
	goji.Serve()
}
