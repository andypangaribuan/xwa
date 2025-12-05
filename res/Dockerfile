ARG GO_VERSION=1.25.5-alpine3.23

FROM golang:${GO_VERSION} AS builder

WORKDIR /goapp
COPY . .

# Install build dependencies for CGO
RUN apk add --no-cache gcc musl-dev

RUN go mod tidy
RUN go mod vendor
RUN CGO_ENABLED=1 \
    GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    go build -mod vendor -ldflags '-extldflags "-static"' -o /.goapp


FROM alpine AS curlstage
RUN apk add --no-cache wget
RUN wget -O /tmp/curl.tar.xz \
    https://github.com/stunnel/static-curl/releases/download/8.17.0-ech/curl-linux-x86_64-musl-8.17.0-ech.tar.xz \
    && tar -xJf /tmp/curl.tar.xz -C /tmp \
    && cp /tmp/curl /usr/bin/curl \
    && chmod +x /usr/bin/curl


FROM gcr.io/distroless/static
LABEL maintainer="Andy Pangaribuan <iam.pangaribuan@gmail.com>"

ARG TZ=Asia/Jakarta
ENV TZ=${TZ}

ARG APP_VERSION=unknown
ENV APP_VERSION=${APP_VERSION}

COPY --from=curlstage /usr/bin/curl /usr/bin/curl
COPY --from=builder /.goapp /goapp
CMD ["/goapp"]
