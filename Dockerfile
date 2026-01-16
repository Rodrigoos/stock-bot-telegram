FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod tidy

COPY . ./

CMD ["go", "run", "./cmd/bot"]
