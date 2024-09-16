ARG GO_VERSION
FROM golang:${GO_VERSION}-alpine AS builder
RUN apk add --no-cache ca-certificates make git curl gcc libc-dev \
    && mkdir -p /build
WORKDIR /build
COPY . /build/
RUN go mod download \
    && make build-linux

FROM golang:${GO_VERSION}-alpine 
RUN apk add --no-cache ca-certificates bash git gcc libc-dev openssh
ENV GO111MODULE=on
COPY --chmod=755 --from=builder /build/supply-chain-linters /bin/supply-chain-linters
COPY --chmod=755 entrypoint.sh /bin/entrypoint.sh
RUN chmod +x /bin/entrypoint.sh

ENTRYPOINT ["/bin/entrypoint.sh"]
