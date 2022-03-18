#docker makefile to build & run

all: build run

build:
	sudo docker build -t app .

run:
	sudo docker run -p 8000:8001 -it app
