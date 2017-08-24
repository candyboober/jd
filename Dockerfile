FROM golang

ADD . /go/src/jd

WORKDIR /go/src/jd

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

RUN go install
