# syntax=docker/dockerfile:1
#
# Uses ./vendor so the image builds without fetching modules (no proxy.golang.org TLS in the container).
# Fixes "tls: bad record MAC" for many VPN / corporate SSL inspection / Docker Desktop network setups.
# After changing go.mod: go mod tidy && go mod vendor

FROM golang:1.22-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY vendor ./vendor
COPY . .

ARG VERSION=dev
ARG COMMIT=none
ARG BUILDTIME=unknown

RUN CGO_ENABLED=0 go build -trimpath -mod=vendor -ldflags="-s -w \
  -X finoraai/backend/internal/version.Version=${VERSION} \
  -X finoraai/backend/internal/version.Commit=${COMMIT} \
  -X finoraai/backend/internal/version.BuildTime=${BUILDTIME}" \
  -o /server ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /server /server

EXPOSE 50051

ENV GRPC_PORT=50051

USER nonroot:nonroot

ENTRYPOINT ["/server"]
