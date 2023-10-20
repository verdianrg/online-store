.PHONY: validate generate build

all: generate build

validate:
	swagger validate ./api/swagger.yml

clean:
	rm -rf server
	go clean -i .

build: clean
	go build -v -installsuffix cgo ./cmd/server

run:
	./server --port=8080 --host=0.0.0.0

generate: validate
	go generate ./...