FROM golang:1.23.1-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./db/models/ ./db/models/
COPY ./db/migrations/ ./db/migrations/
COPY ./db/seed/ ./db/seed/

COPY ./main-service/api/ ./api/
COPY ./main-service/cmd/ ./cmd/
COPY ./main-service/internal/ ./internal/

COPY ./event-service/api/ ./api/
COPY ./event-service/cmd/ ./cmd/
COPY ./event-service/internal/ ./internal/

COPY ./pkg/ ./pkg/

RUN CGO_ENABLED=0 GOOS=linux go build -o /backend-container ./cmd/main.go

EXPOSE 8080

CMD ["/backend-container", "sh"]