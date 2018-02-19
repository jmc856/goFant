FROM golang:latest

ENV DIR_PATH "/go/src"
ADD . $DIR_PATH/gofant
WORKDIR $DIR_PATH/gofant

RUN apt-get update && apt-get install -y \
  ca-certificates \
  libgmp-dev \
  libpq-dev \
  postgresql-client

RUN go get gofant
RUN go build

EXPOSE 8080
