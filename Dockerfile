FROM golang:1.14

WORKDIR /go/src/meerkatonair
COPY . /go/src/meerkatonair

#RUN go get go.mongodb.org/mongo-driver
RUN go get -v -u go.mongodb.org/mongo-driver/mongo
RUN go get github.com/labstack/echo

CMD go run main.go
