all: build run

build:
	docker build -t my-server .

run:
	docker run -p 8080:8080 my-server