test:
	@echo "Running tests"
	go test ./...

run:
	@echo "Starting application"
	go run app.go

format:
	@echo "Formatting the code"
	go fmt

build:
	@echo "Building application"
	go build -o build/stringinator

test-build: test build