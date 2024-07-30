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

COPY --from=builder /build/configs/prod_web.yaml /app/configs/prod_web.yaml

COPY --from=builder /build/configs/jwt.yaml /app/configs/jwt.yaml

CMD ["./web", "-config-path",  "configs/prod_web.yaml"]