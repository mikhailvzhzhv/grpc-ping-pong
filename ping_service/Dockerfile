FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ping_service cmd/ping_service/main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ping_service /app/

CMD [ "./ping_service" ]