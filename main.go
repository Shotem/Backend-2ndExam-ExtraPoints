package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	r := e.Group("/api")
	r.GET("/get", func(c echo.Context) error {
		return c.String(http.StatusOK, "Get Realizado Correctamente")
	})

	r.POST("/post", func(c echo.Context) error {

		value := c.FormValue("Value")
		return c.String(http.StatusAccepted, "Post Realizado Correctamente con valor "+value)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
