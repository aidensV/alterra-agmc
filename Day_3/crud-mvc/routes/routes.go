package routes

import (
	"crud-mvc/constants"
	"crud-mvc/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	r := e.Group("/")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	// unathorize
	e.GET("/books", controllers.GetBookControllers)
	e.POST("/login", controllers.LoginUserController)
	// authorize
	r.POST("/books", controllers.CreateBookControllers)
	r.PUT("/books/:id", controllers.UpdateBookControllers)
	r.DELETE("/books/:id", controllers.DeleteBookControllers)

	r.GET("users/:id", controllers.GetUserDetailController)

	r.GET("/users", controllers.GetUserControllers)
	e.POST("/users", controllers.CreateUserControllers)
	r.PUT("/users/:id", controllers.UpdateUserControllers)
	r.DELETE("/users/:id", controllers.DeleteUserControllers)

	return e
}
