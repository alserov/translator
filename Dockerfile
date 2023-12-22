FROM alpine as ca-certs

RUN apk add -U --no-cache ca-certificates

FROM golang:alpine as builder

WORKDIR /build

COPY ./cmd/translator/main.go ./
COPY ./config/local.yaml ./config/
COPY go.mod go.sum ./
COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin

FROM scratch

COPY --from=builder /build/bin /bin
COPY --from=builder /build/config /config/
COPY --from=ca-certs etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["/bin"]
