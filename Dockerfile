FROM golang:1.25.1-alpine3.22 AS builder

RUN apk add --no-cache git gmp-dev build-base g++ openssl-dev
ADD . /aerium

# Building aerium-daemon
RUN cd /aerium && \
    CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath -o ./build/aerium-daemon ./cmd/daemon && \
    CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath -o ./build/aerium-wallet ./cmd/wallet && \
    CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath -o ./build/aerium-shell ./cmd/shell

## Copy binary files from builder into second container
FROM alpine:3.19

COPY --from=builder /aerium/build/aerium-daemon /usr/bin
COPY --from=builder /aerium/build/aerium-wallet /usr/bin
COPY --from=builder /aerium/build/aerium-shell /usr/bin

ENV WORKING_DIR="/aerium"

VOLUME $WORKING_DIR
WORKDIR $WORKING_DIR
