package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.SetLevel(log.DEBUG)

	// envVars := make(map[string]string)

	// Initialize handler
	// problems := &problems.Handler{
	// 	Env: envVars,
	// }

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "yep we exist"})
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/cyoa", func(c echo.Context) error {
		// problems.Logger = c
		// // return c.HTML(http.StatusOK, "Hello, Docker! <3")
		// // input := []int{3, 5, 13, 14, 5, 12}
		// returnString := problems.ConcurrentTree(10)
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
