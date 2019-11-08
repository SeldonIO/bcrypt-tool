VERSION=$(shell cat ./VERSION)
build: fmt imports lint
	go build -v -ldflags="-w -X main.Version=$(VERSION)" -o bcrypt-tool ./cmd/bcrypt-tool.go

fmt:
	go fmt ./...

imports:
	goimports -w ./

clean:
	rm -f ./bcrypt-tool

lint:
	golint ./...

