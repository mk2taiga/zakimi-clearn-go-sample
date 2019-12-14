FROM golang:1.12.9
LABEL maintainer="zakimi-sample"
WORKDIR /go/src/app
ENV GO111MODULE=on
RUN go mod download
EXPOSE 8080
CMD ["go", "run", "server.go"]