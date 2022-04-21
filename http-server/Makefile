#docker makefile to build & run

all: build run

build:
	docker build -t app .

run:
	docker run -p 8080:8080 -it --rm app

runbg:
	docker run -d -p 8080:8080 --rm app
