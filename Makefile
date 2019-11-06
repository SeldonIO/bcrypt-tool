build: fmt imports lint
	go build -v -o bcrypt-tool ./cmd/bcrypt-tool.go

fmt:
	go fmt ./...

imports:
	goimports -w ./

clean:
	rm -f ./bcrypt-tool

lint:
	golint ./...

