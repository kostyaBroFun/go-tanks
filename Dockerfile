FROM golang:1.19-alpine as builder

COPY . /app
WORKDIR /app
