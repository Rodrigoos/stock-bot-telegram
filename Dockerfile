FROM golang:1.23

RUN apt-get update && apt-get install -y fonts-dejavu-core

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

CMD ["go", "run", "./cmd/bot"]
