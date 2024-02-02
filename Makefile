# Build target
build:
	go build ./cmd/main.go 

# Clean target
clean:
	go clean
	rm -f dontpad-storage

# Test target
test:
	go test -v ./...
