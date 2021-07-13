
all:
	echo hello

validate:
	./hack/update-gofmt.sh
	./hack/verify-gofmt.sh
	./hack/verify-golangci.sh
	go mod tidy
