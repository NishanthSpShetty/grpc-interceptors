GO_VERSION=1.25
ALL_GO_FILES=$(shell find . -type f  -name '*.go')

tidy:
	go mod tidy -compat=$(GO_VERSION)

test:
	go test ./...


fmt:
	gofumpt -w $(ALL_GO_FILES)

