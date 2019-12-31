FROM golang:alpine

MAINTAINER Rémy DEME <demeremy@gmail.com>

RUN apk update && apk add --no-cache git  && \
        go get "github.com/Remydeme/iurgence"

WORKDIR /go/src/github.com/Remydeme/iurgence

EXPOSE 8080

CMD ["make", "build"]
