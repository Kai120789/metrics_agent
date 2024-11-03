FROM golang:1.23.1

WORKDIR /app

COPY ../go.mod ../go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./agent cmd/agent/agent.go

CMD ["./agent"]