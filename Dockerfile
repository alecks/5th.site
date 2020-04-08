FROM golang:alpine

WORKDIR /go/src/5G
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["5g"]