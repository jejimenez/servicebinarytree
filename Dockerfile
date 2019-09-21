FROM golang:1.13.0-alpine3.10
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build cmd/binarytree/main.go
CMD ["/app/main"]