package routes

import (
	"crud-mvc/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/books", controllers.GetBookControllers)
	e.POST("/books", controllers.CreateBookControllers)
	e.PUT("/books/:id", controllers.UpdateBookControllers)
	e.DELETE("/books/:id", controllers.DeleteBookControllers)

	e.GET("/users", controllers.GetUserControllers)
	e.POST("/users", controllers.CreateUserControllers)
	e.PUT("/users/:id", controllers.UpdateUserControllers)
	e.DELETE("/users/:id", controllers.DeleteUserControllers)

	return e
}
