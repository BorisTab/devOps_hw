FROM --platform=linux/amd64 golang:1.19.1-alpine3.16 as builder

COPY . /build

WORKDIR /build

RUN go build -o /tmp/front_server ./main.go

FROM alpine:3.16

COPY --from=builder /tmp/front_server ./front_server
COPY ./static ./static

RUN chmod +x front_server

ENTRYPOINT ["./front_server"]