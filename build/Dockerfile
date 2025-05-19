FROM golang:1.24.2 AS base

FROM base AS builder

WORKDIR /opt/ha-app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./build/app ./cmd/app/main.go

# build small image with execution file
FROM scratch AS dev
WORKDIR /opt/ha-app
ENV CONFIG=./config/local.config.yaml
COPY --from=builder /opt/ha-app/config ./config
COPY --from=builder /opt/ha-app/build/app ./app
EXPOSE 8080
ENTRYPOINT ["./app"]