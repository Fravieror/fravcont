FROM golang:1.8.5-jessie

WORKDIR /go/src/app

ADD src src

CMD ["go", "run", "main.go"]

