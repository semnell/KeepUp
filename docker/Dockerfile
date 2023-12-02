FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/keepup/build/
COPY . .
RUN go get -d -v
RUN go build -o /go/bin/keepup


FROM gcr.io/distroless/static as dev
# Copy our static executable.
WORKDIR /keepup
COPY --from=builder /go/bin/keepup /keepup/keepup
COPY .env config.yaml /keepup/
# Run the hello binary.
ENV GIN_MODE=release
ENTRYPOINT ["/keepup/keepup"]

FROM gcr.io/distroless/static
# Copy our static executable.
WORKDIR /keepup
COPY --from=builder /go/bin/keepup /keepup/keepup
ENV GIN_MODE=release
ENTRYPOINT ["/keepup/keepup"]
