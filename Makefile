PROJECT_NAME := "go-hello"

.PHONY: build

build:
	@GOOS=linux GOARCH=amd64 go build -o bin/$(PROJECT_NAME)_linux_amd64 -i -v $(CMD)
	@GOOS=darwin GOARCH=amd64 go build -o bin/$(PROJECT_NAME)_darwin_amd64 -i -v $(CMD)
	@GOOS=windows GOARCH=amd64 go build -o bin/$(PROJECT_NAME).exe -i -v $(CMD)
