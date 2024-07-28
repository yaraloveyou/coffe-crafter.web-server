FROM golang:1.22.1 AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -v -o /app/web cmd/web/main.go

FROM alpine

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY --from=builder /app/web /app/web

COPY --from=builder /build/configs/web.toml /app/configs/web.toml

CMD ["./web", "-config-path",  "configs/web.toml"]