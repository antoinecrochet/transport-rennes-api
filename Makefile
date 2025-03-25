# Build the project
build:
	go build -o main ./cmd/api/main.go

# Start the application if already built
run:
	./main

# Build the application and start it
start: build run

# Execute tests
test:
	go test ./...