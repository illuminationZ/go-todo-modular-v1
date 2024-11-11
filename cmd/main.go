package main

import (
	"fmt"
	"go-todo-modular-v1/config"
	"go-todo-modular-v1/internal/database"
	"go-todo-modular-v1/internal/server"
	"go-todo-modular-v1/internal/todo"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Build the DSN using the configuration struct
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)

	// Connect to the database
	database.Connect(dsn)

	// Automatically migrate the schema, to keep it up to date
	database.DB.AutoMigrate(&todo.Todo{})

	// Start the server
	server.Run()
}
