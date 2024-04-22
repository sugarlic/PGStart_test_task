all: build_server run_server

build_server:
	@go build -o app ./cmd/web 

run_server:
	@./app

run_server_d:
	@./app &

build_d:
	# docker build -t my-server .
	docker-compose build

run_d:
	# docker run -p 8080:8080 my-server
	docker-compose up

tests:
	go test -v ./...

clean:
	rm app