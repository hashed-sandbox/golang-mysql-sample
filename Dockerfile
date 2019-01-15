FROM golang:1.11.4

RUN go get github.com/oxequa/realize

WORKDIR $GOPATH/src/github.com/hashed-sandbox/golang-mysql-sample
