FROM golang:1.24-alpine

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /service ./cmd/main.go

EXPOSE 8080

# Run
CMD ["/service"]