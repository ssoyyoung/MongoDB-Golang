FROM golang:1.14

WORKDIR /go
COPY . /go

RUN go get go.mongodb.org/mongo-driver

CMD go run main.go