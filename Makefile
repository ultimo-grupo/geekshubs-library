run:
	 cd app && GIN_MODE=release go run main.go

test:
	go get github.com/mfridman/tparse
	cd app && CGO_ENABLED=0 go test ./... -json -cover | tparse -all

build:
	# MacOS
	cd app && GOOS=darwin GOARCH=amd64 go build -o ../bin/main-darwin-amd64 main.go
	# Linux
	cd app && GOOS=linux GOARCH=amd64 go build -o ../bin/main-linux-amd64 main.go
	# Windows
	cd app && GOOS=windows GOARCH=amd64 go build -o ../bin/main-windows-amd64 main.go
