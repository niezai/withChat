package main

import (
	"withChat/internal/config"
	"withChat/internal/routers"
)

func main() {
	config.Init()

	r := routers.NewRouter()

	r.Run(":" + config.Config.Server.Port)
}
