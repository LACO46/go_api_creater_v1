FROM golang:latest

WORKDIR /home

RUN go get github.com/ant0ine/go-json-rest/rest