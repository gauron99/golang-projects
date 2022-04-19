FROM golang AS builder

COPY . /build
WORKDIR /build
RUN go build -o main .

FROM ubuntu
COPY --from=builder /build/main /main
COPY --from=builder /build/env /env
CMD ["./main"]
