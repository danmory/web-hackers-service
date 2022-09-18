FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o go_server github.com/danmory/web-hackers-service/cmd/app

EXPOSE 8010

ENTRYPOINT ["./go_server"]
