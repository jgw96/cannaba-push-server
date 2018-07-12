FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENV GIN_MODE=release

EXPOSE 8080

CMD ["app"]