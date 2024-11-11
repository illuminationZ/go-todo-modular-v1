package server

import (
	"go-todo-modular-v1/internal/database"
	"go-todo-modular-v1/internal/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())  // Logs all requests
	e.Use(middleware.Recover()) // Recovers from panics
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
		},
	)) // Enable CORS

	// Register routes
	router.RegisterRoutes(e, database.DB)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
