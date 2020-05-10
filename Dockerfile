FROM golang:alpine

ENV PORT 8080

RUN apk add --no-cache git \
 && mkdir -p /talks \
 && go get -u golang.org/x/tools/cmd/present

ADD . /talks
WORKDIR /talks

CMD present -play=false -http 0.0.0.0:$PORT
