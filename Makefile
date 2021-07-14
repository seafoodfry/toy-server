
all: run

run:
	docker build -t toy-server .
	docker run -p 8080:8080 toy-server

validate:
	./hack/update-gofmt.sh
	./hack/verify-gofmt.sh
	./hack/verify-golangci.sh
	go mod tidy
	go test ./...
