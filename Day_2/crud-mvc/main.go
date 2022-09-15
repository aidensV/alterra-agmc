package main

import (
	"crud-mvc/config"
	"crud-mvc/routes"

	"github.com/joho/godotenv"
)

func main() {
	config.InitDB()
	e := routes.New()
	godotenv.Load()
	e.Logger.Fatal(e.Start(":8000"))
	// routers.Api()
}
