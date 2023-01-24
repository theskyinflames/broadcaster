FROM golang:1.19.3-alpine3.16

WORKDIR /core-tech

COPY . .

RUN go build cmd/main.go

FROM alpine:3.14.0

WORKDIR /core-tech

COPY --from=0 /core-tech/main .
EXPOSE 80
ENTRYPOINT [ "./main" ]



