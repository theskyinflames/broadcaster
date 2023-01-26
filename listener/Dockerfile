FROM golang:1.19.3-alpine3.16

WORKDIR /listener

COPY . .

RUN go build cmd/main.go

FROM alpine:3.14.0

WORKDIR /listener

COPY --from=0 /listener/main .
EXPOSE 80
ENTRYPOINT [ "./main" ]



