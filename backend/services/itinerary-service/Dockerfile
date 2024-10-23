FROM golang:1.23.2-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY ./api/ ./api/
COPY ./cmd/ ./cmd/
COPY ./internal/ ./internal/

COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

RUN CGO_ENABLED=0 GOOS=linux go build -o /itinerary-service-container ./cmd/main.go

Expose 8083


ENTRYPOINT ["/itinerary-service-container"]