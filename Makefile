run:
	@go version
	@go run main.go
.PHONY: run

test:
	@go version
	@go vet ./...
	@go test -v -timeout 30s -coverprofile=coverage.txt -covermode=atomic ./...
.PHONY: test
