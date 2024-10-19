FROM golang:1.23.1-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY ./cmd/ ./cmd/

RUN CGO_ENABLED=0 GOOS=linux go build -o /itinerary-service-container ./cmd/main.go

Expose 8083


ENTRYPOINT ["/itinerary-service-container"]