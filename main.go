package main

import (
	"main/config"
	"main/routes"
	"os"
)

func main() {
	// config.LoadEnv()
	config.InitDb()
	e := routes.New()
	e.Logger.Fatal(e.Start(envPortOr("8000")))
}

func envPortOr(port string) string {
	envPort := os.Getenv("PORT")
	if envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}
