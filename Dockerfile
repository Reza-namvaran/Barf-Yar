FROM golang:1.24.5 AS builder
WORKDIR /app

COPY panel/go.mod panel/go.sum ./
RUN go mod download

COPY panel/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -o admin-panel ./cmd

RUN CGO_ENABLED=0 GOOS=linux go build -a -o create_admin ./scripts/create_admin.go


FROM debian:bullseye-slim AS runner
WORKDIR /app

COPY --from=builder /app/admin-panel /app/admin-panel
COPY --from=builder /app/create_admin /app/create_admin

COPY panel/static /app/static
COPY panel/internal/templates /app/templates

RUN chmod +x /app/admin-panel /app/create_admin

EXPOSE 8080
ENTRYPOINT ["/app/admin-panel"]
