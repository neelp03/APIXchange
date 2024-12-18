FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
