package main

import (
	"main/config"
	"main/routes"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// config.LoadEnv()
	config.InitDb()
	e := routes.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Timeout())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Logger.Fatal(e.Start(":8080"))
}
