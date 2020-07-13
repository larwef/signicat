all: test build-linux build-mac build-windows

run:
	go run cmd/kitsune/main.go

# PHONY used to mitigate conflict with dir name test
.PHONY: test
test:
	go mod tidy
	go fmt ./...
	go vet ./...
	golint ./...
	go test ./...

doc:
	godoc -http=":6060"
