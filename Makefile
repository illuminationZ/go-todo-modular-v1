# Build the application
all: build

build:
	@echo "Building..."

	@go build -o bin/main main.go

run:
	@go run cmd/main.go

migration:
	@go run cmd/main.go migrate
