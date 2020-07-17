FROM golang:alpine
COPY ./api /app/go
WORKDIR /app/go
RUN go build -o main . 
CMD ["/app/go/main"]
