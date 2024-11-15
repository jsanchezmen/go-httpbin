FROM golang:1.22

ENV SERVER_NAME DEFAULT_SERVER
WORKDIR /httpbin-app
COPY /httpbin-app .
RUN go build
ENTRYPOINT [ "./go-httpbin" ]