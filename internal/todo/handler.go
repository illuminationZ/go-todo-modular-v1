package todo

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateTodoHandler(c echo.Context) error {
	var todo Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	if err := h.service.CreateTodo(&todo); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create todo")
	}
	return c.JSON(http.StatusCreated, todo)
}

func (h *Handler) GetTodosHandler(c echo.Context) error {
	todos, err := h.service.GetAllTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get todos")
	}
	return c.JSON(http.StatusOK, todos)
}

func (h *Handler) GetTodoByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	todo, err := h.service.GetTodoByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Todo not found")
	}
	return c.JSON(http.StatusOK, todo)
}

func (h *Handler) UpdateTodoByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	todo, err := h.service.GetTodoByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Todo not found")
	}

	var updatedTodo Todo
	if err := c.Bind(&updatedTodo); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	// Set ID explicitly to ensure the correct record is updated
	updatedTodo.ID = todo.ID

	// Update only the fields that are provided in the request
	if updatedTodo.Title == "" {
		updatedTodo.Title = todo.Title
	}
	if updatedTodo.Description == "" {
		updatedTodo.Description = todo.Description
	}
	if !updatedTodo.Completed {
		updatedTodo.Completed = todo.Completed
	}

	if err := h.service.UpdateTodoByID(&updatedTodo); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update todo")
	}

	return c.JSON(http.StatusOK, "Todo updated successfully")
}

func (h *Handler) DeleteTodoByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	if err := h.service.DeleteTodoByID(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete todo")
	}

	return c.NoContent(http.StatusNoContent)
}
