FROM golang:1.15 as build-env

WORKDIR /go/src/app
ADD go.mod go.sum /go/src/app/
RUN go mod download


ADD cmd/ /go/src/app/cmd/
ADD pkg/ /go/src/app/pkg/
WORKDIR /go/src/app/cmd
RUN go build -o /go/bin/app


FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/app /
CMD ["/app"]
