FROM golang:alpine
COPY ./api /app/go
WORKDIR /app/go
RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get github.com/oxequa/realize \
    && go build -o main.
EXPOSE 8000
CMD ["/app/go/main"]
