all: build_server run_server

build_server:
	@go build -o app ./cmd/web 

run_server:
	@./app

build_d:
	docker build -t my-server .

run_d:
	docker run -p 8080:8080 my-server

clean:
	rm app