FROM golang:1.14

WORKDIR /app
COPY ./ /app

RUN go mod download
RUN go mod vendor
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main