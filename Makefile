build:
	go build -o bin/main main.go

run:
	 GIN_MODE=release go run main.go