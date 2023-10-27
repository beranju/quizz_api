package main

import (
	"main/config"
	"main/routes"
)

func main() {
	config.LoadEnv()
	config.InitDb()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
