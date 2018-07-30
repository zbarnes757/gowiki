FROM golang:1.10

WORKDIR /go/src/app
COPY . .

EXPOSE 3000

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
