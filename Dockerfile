# syntax = docker/dockerfile:experimental
# Build Container
FROM golang:1.19 as builder

ENV GO111MODULE on
ENV GOPRIVATE=github.com/latonaio
WORKDIR /go/src/github.com/latonaio

COPY . .
RUN go mod download
RUN go build -o convert-to-jp-pint-peppol-from-dpfm-invoice-document ./

# Runtime Container
FROM alpine:3.14
RUN apk add --no-cache libc6-compat
ENV SERVICE=convert-to-jp-pint-peppol-from-dpfm-invoice-document \
    APP_DIR="${AION_HOME}/${POSITION}/${SERVICE}"

WORKDIR ${AION_HOME}

COPY --from=builder /go/src/github.com/latonaio/convert-to-jp-pint-peppol-from-dpfm-invoice-document .

CMD ["./convert-to-jp-pint-peppol-from-dpfm-invoice-document"]
