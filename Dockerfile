FROM golang:1.23.1-alpine AS builder
RUN apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN go build -o /goapp main.go
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:latest
RUN apk --no-cache add ca-certificates bash postgresql-client

COPY --from=builder /goapp /goapp
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY ./migrations /app/migrations

ENV CONFIG_PATH="/config/local.yaml"

EXPOSE 8080


COPY ./entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh


CMD ["/entrypoint.sh"]
