FROM golang:alpine as base
RUN apk add --no-cache git

WORKDIR /go/src/app

COPY main.go .
RUN go get -d -v ./...


FROM base
RUN go install -v ./...

CMD ["app"]
