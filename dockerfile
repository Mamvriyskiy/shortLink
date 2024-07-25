FROM golang:latest

COPY ./ ./
RUN go build -o main main.go
CMD ["./main"]
