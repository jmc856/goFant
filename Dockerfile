FROM golang:latest 
ENV PACKAGE_PATH="$GOPATH/bitbucket.org/jcalvert55/gofant/src"
RUN mkdir -pv $PACKAGE_PATH 
ADD . $PACKAGE_PATH 
WORKDIR $PACKAGE_PATH
RUN go get -v bitbucket.org/jcalvert55/gofant
RUN go run main.go 
