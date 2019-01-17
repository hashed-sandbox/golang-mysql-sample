FROM golang:1.11.4

WORKDIR $GOPATH/src/github.com/hashed-sandbox/golang-mysql-sample

COPY . .

RUN \
  curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
  dep ensure && \
  go get github.com/oxequa/realize
