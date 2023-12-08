package main

import (
	"withChat/config"
	"withChat/internal/db"
	"withChat/internal/router"
)

func main() {
	config.Init()
	db.Initialize()

	r := router.NewRouter()

	r.Run(":" + config.Config.Server.Port)
}
