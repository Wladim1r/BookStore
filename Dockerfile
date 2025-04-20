### Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.sum go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/lib ./cmd/app/main.go

### Final stage
FROM alpine:3.21

WORKDIR /book

COPY --from=builder /app/bin/lib .
COPY --from=builder /app/internal/database /book/internal/database

CMD [ "./lib" ]