FROM golang:1.23.1-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY ./api/ ./api/
COPY ./cmd/ ./cmd/
COPY ./internal/ ./internal/
COPY ./utils/ ./utils/

RUN CGO_ENABLED=0 GOOS=linux go build -o /itinerary-service-container ./cmd/main.go

EXPOSE 8081

ENTRYPOINT ["/itinerary-service-container"]