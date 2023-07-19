# from golang alpine image
FROM golang:1.20-alpine as builder

# labels for swarm and other managers
LABEL app="automated-pen-testing-runner"
LABEL docker_file_version="v0.0.1"
LABEL database_driver="mysql"

# maintainer
MAINTAINER amirhossein.najafizadeh21@gmail.com

# app work directory
WORKDIR /app

# env variables
ENV VERSION="v0.0.1"
ENV APP="automated-pen-testing-runner"

# copy go.mod and go.sum
COPY go.mod go.sum ./

# download deps
RUN go mod download

# copy all files
COPY . .

# building go executable file
RUN go build -o main

# second stage
FROM alpine

# create src work directory
WORKDIR /src

# copy execute file
COPY --from=builder /app/main main

# expose port 8080
EXPOSE 8080

# start http service
CMD ./main --port 8080 --config "config.yml"
