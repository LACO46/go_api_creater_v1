FROM golang:latest

WORKDIR /home

RUN go get -u github.com/labstack/echo/...