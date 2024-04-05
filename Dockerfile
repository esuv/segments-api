# Image page: <https://hub.docker.com/_/golang>
FROM golang:1.21-alpine AS builder

RUN apk add --no-cache ca-certificates

WORKDIR /usr/local/src

COPY ["./go.mod", "./go.sum", "./"]

# Burn dependencies cache
RUN set -x \
    && go version \
    && go mod download \
    && go mod verify

COPY . ./
RUN go build -o /tmp/app ./cmd/app/main.go

# Image page: <https://hub.docker.com/_/alpine>
FROM alpine:latest AS runtime

COPY --from=builder /tmp/app /bin/app

ENTRYPOINT ["/bin/app"]
