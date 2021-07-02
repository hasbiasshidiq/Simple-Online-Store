# Builder
FROM golang:1.16.5-alpine3.14 as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY . .

RUN go build -o store-api.exe

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 8888

COPY --from=builder /app/store-api.exe /app

CMD /app/store-api.exe