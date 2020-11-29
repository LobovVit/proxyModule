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
    mv instantclient_19_9/* /lib && \
    rm -rf ora_client_x64-19.9.zip
# Link Libs
RUN ln -sf /lib/libclntsh.so.19.1 /lib/libclntsh.so;
RUN cd /lib && \
 ln -s /lib64/* /lib && \
 ln -s /usr/lib/libnsl.so.2 /usr/lib/libnsl.so.1 && \
 ln -s /usr/lib/libc.so /usr/lib/libresolv.so.2
# The libnsl version installed by Alpine is the number two, and instaclient use the number 1 so, create symbolic link for number one
RUN ln -sf /usr/lib/libnsl.so.2.0.0  /usr/lib/libnsl.so.1
WORKDIR /root/
COPY --from=builder /app/proxyModule ./
CMD ["./proxyModule"]