FROM golang:1.7
RUN go get github.com/astaxie/beego && \
    go get github.com/beego/bee

EXPOSE 8080

CMD ["bee", "run"]
