# Builder stage
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/keepup/build/
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /go/bin/keepup

# Dev stage
FROM alpine:3.18.5 as dev
WORKDIR /keepup
COPY --from=builder /go/bin/keepup /keepup/keepup
COPY .env config.yaml /keepup/
ENV FAKTORY_PROVIDER=KEEPUP_FAKTORY_URL
ENTRYPOINT ["/keepup/keepup"]

# Production stage
FROM gcr.io/distroless/static:nonroot as prod
WORKDIR /keepup
COPY --from=builder /go/bin/keepup /keepup/keepup
ENV GIN_MODE=release
ENV FAKTORY_PROVIDER=KEEPUP_FAKTORY_URL
ENTRYPOINT ["/keepup/keepup"]
