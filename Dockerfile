FROM golang:latest
ENV DIR_PATH "/go/src"
ADD . $DIR_PATH/gofant
WORKDIR $DIR_PATH/gofant
RUN go get github.com/lib/pq
RUN go get gofant
