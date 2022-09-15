package controllers

import (
	"crud-mvc/lib/database"
	"crud-mvc/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetUserControllers(c echo.Context) error {
	users, e := database.GetUsers()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success", "users": users,
	})

}

func CreateUserControllers(c echo.Context) error {
	var users models.User
	err := c.Bind(&users)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	_, e := database.CreateUser(users)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success", "users": users,
	})

}

func UpdateUserControllers(c echo.Context) error {
	var users models.User
	err := c.Bind(&users)

	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	users.ID = uint(id)
	_, e := database.UpdateUser(users)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success", "users": users,
	})

}

func DeleteUserControllers(c echo.Context) error {
	var users models.User
	err := c.Bind(&users)

	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	users.ID = uint(id)
	_, e := database.DeleteUser(users)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success", "users": users,
	})

}
