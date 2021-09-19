# ==============================================================================
# Main

run:
	go run ./cmd/rest-api

build:
	go build ./cmd/rest-api

test:
	go test -cover ./...

fmt:
	go fmt ./...

mock:
	moq -rm -out internal/repository/book_respository_mock.go internal/core/port BookRepository

# ==============================================================================
# Docker support

docker-build:
	docker build -t rest-api -f docker/Dockerfile .

.PHONY: run build test fmt mock docker-build
