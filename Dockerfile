FROM golang:1.14

WORKDIR /go/src/github.com/ssoyyoung.p/MongoDB-Golang
#COPY . /go/src/github.com/ssoyyoung.p/MongoDB-Golang

#RUN go get go.mongodb.org/mongo-driver
RUN go get -v -u go.mongodb.org/mongo-driver/mongo
RUN go get github.com/labstack/echo

CMD go run main.go
