FROM golang AS builder

RUN mkdir -p /build
RUN echo "buidler"
RUN ls
ADD . /build

WORKDIR /build
RUN go build -o main .

FROM alpine
RUN echo "alpine"
RUN ls
COPY --from=builder 
CMD ["/build/main"]
