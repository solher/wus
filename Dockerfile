FROM golang

MAINTAINER solher

EXPOSE 3000

RUN go get github.com/tools/godep

WORKDIR /go/src/github.com/solher/wus

CMD godep restore && go build -v && /go/src/github.com/solher/wus
