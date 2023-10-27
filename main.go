package main

import (
	"main/config"
	"main/routes"
	"time"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// config.LoadEnv()
	config.InitDb()
	e := routes.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Logger.Fatal(e.Start(":8080"))
}
