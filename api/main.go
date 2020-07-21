package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
)

func main() {
    e := echo.New()
	e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello motio")
    })
    e.Logger.Fatal(e.Start(":8000"))
}

func echoHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello motio</h1>")
}
