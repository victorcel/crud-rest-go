FROM golang:1.17

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o app .

EXPOSE 5050
CMD ["./app"]
