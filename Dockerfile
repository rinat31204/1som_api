# ---------- build stage ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# зависимости
COPY go.mod go.sum ./
RUN go mod download

# исходники
COPY . .

# сборка
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# ---------- runtime stage ----------
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["/app/app"]
