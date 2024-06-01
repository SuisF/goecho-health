package main

import (
	"github.com/labstack/echo"
	"github.com/SuisF/goecho-health/cmd/api/handlers"
	)

func main() {
	 e := echo.New()
	 e.GET("/health-check", handlers.HealthCheckHandler)
	 e.GET("/posts", handlers.PostIndexHandler)
	 e.GET("/post/:id", handlers.PostSingleHandler)

	 e.Logger.Fatal(e.Start(":2500"))
}

