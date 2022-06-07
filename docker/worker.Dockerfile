FROM golang:1.16-buster AS build

WORKDIR /go/src/lstrgiang/gitscanner

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/worker ./cmd/worker


###

FROM alpine:3.9

COPY --from=0 /usr/local/bin/worker /usr/local/bin/worker
RUN apk add --no-cache ca-certificates git

ENTRYPOINT ["worker"]

