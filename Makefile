build:
	cd app && go build -o ../bin/main main.go

run:
	 cd app && GIN_MODE=release go run main.go

test:
	cd app && CGO_ENABLED=0 go test ./...