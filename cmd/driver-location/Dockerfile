FROM golang:1.23.0 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o driver-location ./cmd/driver-location/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/driver-location .
COPY --from=builder /app/config/driver-location.yaml ./config/driver-location.yaml

CMD ["./driver-location"]