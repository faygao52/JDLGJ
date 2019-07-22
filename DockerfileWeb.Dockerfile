FROM golang:1.12 AS build

MAINTAINER Jingyi Gao <faygao52@gmail.com>

# Checkout our code onto the Docker container
WORKDIR /app


COPY go.mod .
COPY go.sum .
RUN go mod download

ADD . /app

RUN go build

# Expose a port to run our application
EXPOSE 8000

# Run the server command
ADD runweb.sh run.sh
RUN chmod +x run.sh
CMD ./run.sh