FROM golang:alpine
COPY ./api /app/go
WORKDIR /app/go
RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/labstack/echo/middleware \
    && go get github.com/jinzhu/gorm \
    && go get github.com/oxequa/realize \
    && go build -o main.
EXPOSE 8000
CMD ["/app/go/main"]
