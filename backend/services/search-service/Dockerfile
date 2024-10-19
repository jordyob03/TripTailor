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


RUN CGO_ENABLED=0 GOOS=linux go test ./... -v 
RUN CGO_ENABLED=0 GOOS=linux go build -o /search-service-container ./cmd/main.go

EXPOSE 8084

ENTRYPOINT ["/search-service-container"]
