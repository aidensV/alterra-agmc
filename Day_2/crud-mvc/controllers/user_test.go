package controllers

import (
	"crud-mvc/config"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	config.InitDB()
	e := echo.New()

	return e
}

func TestGetUserControllers(t *testing.T) {

	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "berhasil",
			path:                 "/users",
			expectBodyStartsWith: "{\"status\":\"success\",\"users\":[",
			expectStatus:         http.StatusOK,
		},
		{
			name:                 "error",
			path:                 "/users",
			expectBodyStartsWith: "{\"status\":\"error\",\"users\":[",
			expectStatus:         http.StatusMethodNotAllowed,
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	for _, testCases := range testCases {
		c.SetPath(testCases.path)

		if assert.NoError(t, GetUserControllers(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCases.expectBodyStartsWith))
		}

		if assert.NoError(t, GetUserControllers(c)) {
			assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCases.expectBodyStartsWith))
		}
	}
}
