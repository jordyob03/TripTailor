FROM golang:1.23.1-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./db/models/ ./db/models/
COPY ./db/migrations/ ./db/migrations/
COPY ./db/seed/ ./db/seed/
COPY ./db/tests/ ./db/tests/

COPY ./main-service/api/ ./main-service/api/
COPY ./main-service/cmd/ ./main-service/cmd/
COPY ./main-service/internal/ ./main-service/internal/

COPY ./event-service/api/ ./event-service/api/
COPY ./event-service/cmd/ ./event-service/cmd/
COPY ./event-service/internal/ ./event-service/internal/

COPY ./auth-service/api/ ./auth-service/api/
COPY ./auth-service/cmd/ ./auth-service/cmd/
COPY ./auth-service/internal/ ./auth-service/internal/

COPY ./profile-service/api/ ./profile-service/api/
COPY ./profile-service/cmd/ ./profile-service/cmd/
COPY ./profile-service/internal/ ./profile-service/internal/

COPY ./pkg/ ./pkg/

RUN CGO_ENABLED=0 GOOS=linux go build -o /backend-container ./main-service/cmd/main.go

EXPOSE 8080

CMD ["/backend-container"]
