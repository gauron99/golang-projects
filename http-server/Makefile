#docker makefile to build & run

all: build run

build:
	docker build -t app .

run:
	docker run -p 8000:8000 -it --rm app

runbg:
	docker run -d -p 8000:8000 --rm app
