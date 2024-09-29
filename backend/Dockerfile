## Requires a compose file, in the meantime:
## cd {$PROJECT_ROOT}/backend docker build --name=go-container . 
## docker run -p 8080:8080 go-container 
FROM golang:1.23.1-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY /main-service/cmd/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-container

EXPOSE 8080

CMD ["/go-container"]


