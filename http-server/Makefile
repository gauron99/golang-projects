#docker makefile to build & run

all: build run

build:
	docker build -t app .

run:
	docker run -p 8000:8001 -it app
