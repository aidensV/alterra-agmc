package main

import (
	"crud-mvc/config"
	middle "crud-mvc/middelware"
	"crud-mvc/routes"

	"github.com/joho/godotenv"
)

func main() {
	config.InitDB()
	e := routes.New()
	middle.LogMiddlewares(e)
	godotenv.Load()
	e.Logger.Fatal(e.Start(":8000"))
	// routers.Api()
}
