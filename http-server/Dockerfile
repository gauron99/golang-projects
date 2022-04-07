FROM golang AS builder

RUN mkdir -p /build
ADD . /build

WORKDIR /build
RUN go build -o main .

FROM alpine
COPY --from=builder ./build/main ./main 
CMD ["/build/main"]
