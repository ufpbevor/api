install-ci-dep-test:
	sh test/go.sh
	sh test/mongodb.sh

run:
	./api -api-port=8080 -mongo-url=mongodb://127.0.0.1:27017

build:
	$(GOROOT)/bin/go build -o api cmd/api/main.go

build-linux:
	env GOOS=linux $(GOROOT)/bin/go build -o api cmd/api/main.go

go: build run

test-all:
	$(GOROOT)/bin/go test ./core -v

test-all-cover:
	$(GOROOT)/bin/go test ./core -cover -coverprofile cover.out
	$(GOROOT)/bin/go tool cover -html=cover.out -o cover.html
