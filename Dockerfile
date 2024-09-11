FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod ./
COPY main.go ./

RUN go build -v -o /usr/local/bin/app

CMD ["app"]

