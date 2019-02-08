FROM golang:1.7.3
ENV GOBIN=/go/bin
RUN mkdir -p /go/src/github.com/shorten_url
WORKDIR /go/src/github.com/shorten_url
ADD shorten_url /go/src/github.com/shorten_url/
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/shorten_url/shorten_url .
CMD ["./shorten_url"] 
