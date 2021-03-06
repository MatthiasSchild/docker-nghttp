FROM golang:1.15 as builder
WORKDIR /etc/nghttp
COPY ./app .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o nghttp .

FROM alpine:3.12
ENV WORK_DIR /etc/nghttp/web
WORKDIR /etc/nghttp
COPY --from=builder /etc/nghttp/nghttp /etc/nghttp/nghttp
RUN mkdir /etc/nghttp/web
VOLUME /etc/nghttp/web
EXPOSE 8080
CMD ["/etc/nghttp/nghttp"]
