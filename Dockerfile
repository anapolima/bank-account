FROM golang:1.18.2-alpine

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
RUN mkdir app

COPY . /app
COPY ./go.mod /app
COPY ./go.sum /app
WORKDIR /app
RUN go mod download
RUN go build -o main /app/cmd/main.go

ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

EXPOSE 8080
CMD ["/app/main"]
