FROM golang:1.19.3-alpine3.16

WORKDIR /publisher

COPY . .

RUN go build cmd/main.go

FROM alpine:3.14.0

WORKDIR /publisher

COPY --from=0 /publisher/main .
EXPOSE 80
ENTRYPOINT [ "./main" ]



