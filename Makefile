build:
	@go build -v -o example ./cmd/example

test:
	@go test -v -coverprofile .coverage -race -timeout 30s ./...
	@go tool cover -func .coverage

coverage: test
	@go tool cover -html=.coverage

.DEFAULT_GOAL := build