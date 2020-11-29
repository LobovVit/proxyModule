FROM golang:1.15.5 as builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
COPY main.go ./
RUN go build -o main .
CMD ["/app/main"]

FROM alpine:latest as my_im
# Setup LD lib path
ENV LD_LIBRARY_PATH=/lib
# Install Apline dependencies
RUN apk add gcc libnsl libaio unzip openssl-dev autoconf musl-dev libc6-compat
RUN wget https://download.oracle.com/otn_software/linux/instantclient/199000/instantclient-basic-linux.x64-19.9.0.0.0dbru.zip && \
    unzip instantclient-basic-linux.x64-19.9.0.0.0dbru.zip && \
    cp -r instantclient_19_9/* /lib && \
    rm -rf instantclient-basic-linux.x64-19.9.0.0.0dbru.zip
# Link Libs
RUN ln -sf /lib/libclntsh.so.19.1 /lib/libclntsh.so;
RUN cd /lib && \
 ln -s /lib64/* /lib && \
 ln -s /usr/lib/libnsl.so.2 /usr/lib/libnsl.so.1 && \
 ln -s /usr/lib/libc.so /usr/lib/libresolv.so.2
# The libnsl version installed by Alpine is the number two, and instaclient use the number 1 so, create symbolic link for number one
RUN ln -sf /usr/lib/libnsl.so.2.0.0  /usr/lib/libnsl.so.1
WORKDIR /root/
COPY --from=builder /app/main ./
CMD ["./main"]