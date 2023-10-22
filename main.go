package main

import (
	"main/config"
	"main/routes"
)

func main() {
	config.InitDb()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
