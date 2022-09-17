package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Book struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Isbn string `json:"isbn"`
}

var books = []Book{
	{ID: 1, Name: "Programing golang", Isbn: "9022277888"},
	{ID: 2, Name: "Programing python", Isbn: "9028888008"},
}

func GetBookControllers(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success", "data": books,
	})

}

func CreateBookControllers(c echo.Context) error {
	book := Book{}
	c.Bind(&book)
	fmt.Println(book.Name)
	books = append(books, book)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success", "users": books,
	})

}

func UpdateBookControllers(c echo.Context) error {
	book := Book{}
	c.Bind(&book)
	id := c.Param("id")
	idRef, _ := strconv.ParseInt(id, 10, 64)
	for i := 0; i < len(books); i++ {

		if books[i].ID == idRef {

			books = books[:i+copy(books[i:], books[i+1:])]
		}
	}
	book.ID = idRef
	books = append(books, book)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success", "users": books,
	})

}

func DeleteBookControllers(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	for i := 0; i < len(books); i++ {

		if books[i].ID == id {

			books = books[:i+copy(books[i:], books[i+1:])]
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success", "users": books,
	})

}
