FROM golang:1.13.0-stretch

ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0 PORT=8888

WORKDIR /loaner

COPY . .

EXPOSE 8888

RUN go build loaner/graphql/server;

ENTRYPOINT ["./server"]