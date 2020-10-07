FROM golang:latest

LABEL maintainer="Achilleas Tzenetopoulos <atzenetopoulos@microlab.ntua.gr>"



WORKDIR /go/src/github.com/iwita/node-agent/
COPY server/ /go/src/github.com/iwita/node-agent/server
COPY info /go/src/github.com/iwita/node-agent/info

#RUN go mod download
WORKDIR /go/src/github.com/iwita/node-agent
RUN go mod init node-agent
RUN go build -o server/main server/main.go


FROM ubuntu:18.04
#RUN apk --no-cache add ca-certificates
#RUN apk add util-linux
WORKDIR /
COPY --from=0 /go/src/github.com/iwita/node-agent/ /node-agent
WORKDIR /node-agent/

RUN chmod u+x server/pin.sh
RUN chmod u+x server/getCores.sh
EXPOSE 4242

#CMD ["/server/server/main"]
CMD ["/node-agent/server/main"]