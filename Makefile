.PHONY: dev build

dev:
	@echo "Starting development..."
	npx concurrently -c "#93c5fd,#c4b5fd,#54e87b" "npx tailwindcss -i internal/template/app.css -o pb_public/style.css --watch" "templ generate --watch" "air" --names=tailwindcss,templ,air

build:
	@echo "Building..."
	npx tailwindcss -i internal/template/app.css -o pb_public/style.css
	templ generate
	go build -o bin/http cmd/http/main.go

run:
	@echo "Running..."
	go run cmd/http/main.go serve
