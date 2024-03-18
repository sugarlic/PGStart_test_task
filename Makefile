all: build run

build:
	@go build -o app ./cmd/web 

run:
	@./app

build_d:
	docker build -t my-server .

run_d:
	docker run -p 8080:8080 my-server

clean:
	rm app