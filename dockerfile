
FROM golang:1.23.2-bookworm AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .


FROM alpine:latest


WORKDIR /root/
COPY --from=builder /app/myapp .
EXPOSE 8080

CMD ["./myapp"]
