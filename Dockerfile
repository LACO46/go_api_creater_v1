FROM golang:latest

RUN go get -u github.com/labstack/echo/...
RUN go get -v github.com/rubenv/sql-migrate/...