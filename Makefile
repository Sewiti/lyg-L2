NAME = lyg

build:
	go build -o ./bin/$(NAME) ./cmd/$(NAME)

run:
	go run ./cmd/$(NAME)

compile:
	# Linux 64-bit
	GOOS=linux GOARCH=amd64 go build -o ./bin/$(NAME)-linux-amd64 ./cmd/$(NAME)
	# Windows 64-bit
	GOOS=windows GOARCH=amd64 go build -o ./bin/$(NAME)-windows-amd64.exe ./cmd/$(NAME)
