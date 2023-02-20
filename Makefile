GO_VERSION=1.17

tidy:
	go mod tidy -compat=$(GO_VERSION)

test:
	go test ./...

