FROM golang:1.7
RUN go get github.com/astaxie/beego && \
    go get github.com/beego/bee && \
    go get github.com/goinggo/tracelog && \
    go get github.com/smartystreets/goconvey/convey

EXPOSE 8080

CMD ["bee", "run"]
