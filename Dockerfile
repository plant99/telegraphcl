FROM golang:1.15.4-alpine3.12
RUN apk add git
# create a working directory
ADD . /go/src/github.com/plant99/telegraphcl
WORKDIR /go/src/github.com/plant99/telegraphcl/cmd/telegraphcl

RUN go install .

WORKDIR /root/telegraph-blogs/
# run main.go
ENTRYPOINT [ "telegraphcl" ]