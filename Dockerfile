# Build stage
FROM golang:1.19 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=1 go build -o ./bin/main ./cmd/api/main.go

# Run stage
FROM alpine:3.16

WORKDIR app

RUN adduser dsha256 --disabled-password --no-create-home
RUN chown -R dsha256:dsha256 /app
RUN chmod -R 755 /app

COPY --from=builder ./app/bin/main /app
CMD [ "/app/main" ]

USER dsha256
