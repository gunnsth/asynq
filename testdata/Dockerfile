#!------- builder
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /build

# Copy go mod and sum files
COPY testdata/go.mod .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . .
RUN mkdir -p /mybin

#!------- taskqueue
FROM builder AS taskqueue
RUN go build -o /mybin/taskqueue ./testdata/taskqueue.go

#!------- taskrunner
FROM builder AS taskrunner
RUN go build -o /mybin/taskrunner ./testdata/taskrunner.go

