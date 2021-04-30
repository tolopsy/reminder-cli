.PHONY: client

client:
	@echo "Building the client binary"
	go build -o bin/client.exe cmd/client/main.go