FROM golang:1.5.1
MAINTAINER upccup

ENV GOPATH /go

WORKDIR /go/src/

ADD ./ /go/src/

RUN export GO15VENDOREXPERIMENT="1" && \
    CGO_ENABLED=0 go build -v -o verify && \
    mv verfiy /verify && \

CMD ["/bin/bash", "-c", "while true; do echo hello omega; sleep 1; done"]
