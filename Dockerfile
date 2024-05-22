# Choose whatever you want, version >= 1.16
FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go install -mod=mod github.com/githubnemo/CompileDaemon
RUN go mod tidy

ENTRYPOINT CompileDaemon --build="go build -o build/goapp" -command="./build/goapp" -build-dir=/app -polling=true
