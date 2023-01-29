# Build stage
FROM golang:1.18 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o ./bin/main ./cmd/api/main.go

# Run stage
FROM alpine:3.14
COPY --from=builder ./app/bin/main /
CMD [ "./main" ]    