FROM golang:1.23.1-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY ./api/ ./api/
COPY ./cmd/ ./cmd/
COPY ./internal/ ./internal/
COPY ./utils/ ./utils/

COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-service-container ./cmd/main.go

EXPOSE 8081

ENTRYPOINT ["/auth-service-container"]

