FROM golang:alpine

WORKDIR /app

ADD . .

RUN go mod download

RUN go build -o main .

RUN go get github.com/githubnemo/CompileDaemon

# RUN ["chmod", "+x", "docker/execute.sh”]


ENTRYPOINT ["./main"]