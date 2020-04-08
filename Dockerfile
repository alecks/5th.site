FROM golang:alpine

WORKDIR /go/src/5g
COPY . .

RUN apk add git
RUN go get -v .
RUN go build -v .

ENTRYPOINT ["/go/src/5g/5g"]

EXPOSE 80