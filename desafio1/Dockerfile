FROM golang:1.16-alpine3.14

WORKDIR /go/src

COPY /src .

ENV PATH="/go/bin:${PATH}"

RUN GOOS=linux go build main.go

EXPOSE 8000

CMD ["./main"]