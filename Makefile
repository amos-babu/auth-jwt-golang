BINARY=bin/auth-jwt

build:
	go build -o $(BINARY) cmd/main.go

run:
	go run cmd/main.go

clean:
	rm -rf $(BINARY)
