FROM golang

ADD . /go/src/golangBowlingAPI

RUN go install /go/src/golangBowlingAPI/bowlingServer/main

ENTRYPOINT ["/go/bin/golangBowlingAPI/bowlingServer"]

EXPOSE 5300
