FROM golang:alpine

RUN mkdir /build

ADD . /build

WORKDIR /build

RUN go build -o main .

CMD ["/build/main"]
