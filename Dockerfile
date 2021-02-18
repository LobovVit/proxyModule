FROM golang:1.15.5 as builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
COPY proxyModule.go ./
RUN go build -o proxyModule .
CMD ["/app/proxyModule"]

FROM alpine:latest as proxyModule
# Setup LD lib path
ENV LD_LIBRARY_PATH=/lib
# Install Apline dependencies
RUN apk add gcc libnsl libaio unzip openssl-dev autoconf musl-dev libc6-compat
ADD ora_client_x64-19.9.zip /ora_client_x64-19.9.zip
RUN unzip ora_client_x64-19.9.zip && \
    cp -r instantclient_19_9/* /lib && \
    rm -rf ora_client_x64-19.9.zip && \
    apk add libaio && \
    apk add libaio libnsl libc6-compat

COPY script.sh /root/script.sh
RUN /root/script.sh
WORKDIR /root/
COPY --from=builder /app/proxyModule ./
CMD ["./proxyModule"]