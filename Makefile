all: build_server run_server

build_server:
	@go build -o app ./cmd/web 

run_server:
	@./app

run_server_d:
	@./app &

build_d:
	docker build -t my-server .

run_d:
	docker run -p 8080:8080 my-server

tests: build_server
	bash test.sh
	rm -rf app

clean:
	rm app