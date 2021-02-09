FROM    golang:1.15.8-alpine3.13 as go_builder
RUN     sed -i -e 's/https/http/' /etc/apk/repositories && apk update && apk add --no-cache git
COPY    . /go/src/kafka-producer
ENV     GOPROXY=http://proxy.golang.org
WORKDIR /go/src/kafka-producer
RUN     go get -d -v -insecure . && go build -o kafka-producer .

FROM    alpine:3.13.1
COPY    --from=go_builder   /go/src/kafka-producer .
EXPOSE  8080
CMD     ["./kafka-producer"]