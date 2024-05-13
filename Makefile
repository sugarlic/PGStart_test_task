all: build_d run_d

build_server:
	@go build -o app ./cmd/web 

run_server:
	@./app

run_server_d:
	@./app &

build_d:
	docker compose build

run_d:
	docker compose up -d

stop:
	docker compose stop

tests:
	go test -v ./...

cl_build:
	go build -o cl ./client/client.go

clean:
	rm app