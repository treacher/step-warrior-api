FROM golang

ADD . /go/src/github.com/treacher/step-warrior-api
RUN go install github.com/treacher/step-warrior-api

CMD ["/go/bin/step-warrior-api"]

EXPOSE 8080
