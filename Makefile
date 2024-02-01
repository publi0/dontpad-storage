# Build target
build:
	go build ./cmd/main.go 

# Clean target
clean:
	go clean
	rm -f dontpad-storage

# Test target
test:
	$(GOTEST) -v ./...

# Get dependencies target
get:
	$(GOGET) -v ./...

.PHONY: build clean test get
