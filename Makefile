
test: 
	go test -v -cover -covermode=atomic ./...

engine:
	go build cmd/binarytree/main.go

unittest:
	go test -short  ./...

docker:
	docker build -t servicebinarytreeapp .

run:
	docker run -p 8080:8080 -it servicebinarytreeapp

run-with-data:
	docker run -p 8080:8080 -it servicebinarytreeapp

