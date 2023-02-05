FROM golang:1.16-alpine
WORKDIR /splinter

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go mod tidy

RUN go env -w GO111MODULE=off

RUN go build -o splinter

EXPOSE 8081

ENTRYPOINT [ "go","run","main.go" ]