FROM golang:1.23.0 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o matching-api ./cmd/matching-api/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/matching-api .
COPY --from=builder /app/config/matching-api.yaml ./config/matching-api.yaml

CMD ["./matching-api"]
