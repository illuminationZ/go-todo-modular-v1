package router

import (
	"go-todo-modular-v1/internal/todo"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	repo := todo.NewRepository(db)
	service := todo.NewService(repo)
	handler := todo.NewHandler(service)

	e.POST("/todos", handler.CreateTodoHandler)
	e.GET("/todos", handler.GetTodosHandler)
	e.GET("/todos/:id", handler.GetTodoByIDHandler)
	e.PATCH("/todos/:id", handler.UpdateTodoByIDHandler)
	e.DELETE("/todos/:id", handler.DeleteTodoByIDHandler)
}
