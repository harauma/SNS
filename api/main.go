package main

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo"
)

type User struct {
    Name  string `json:"name"`
    Age int `json:"age"`
}

func main() {
    e := echo.New()
    initRouting(e)
    e.Logger.Fatal(e.Start(":8000"))
}

func initRouting(e *echo.Echo) {
    e.GET("/", echoHello)
    e.GET("/users", getUsers)
	e.GET("/user/:userId", getUser)
    e.POST("/user", createUser)
}

func echoHello(c echo.Context) error {
    fmt.Println("echoHelloが呼ばれました！")
    return c.String(http.StatusOK, "Hello motio")
}

func getUsers(c echo.Context) error {
    fmt.Println("getUsersが呼ばれました！")
    name := c.QueryParam("name")
    age := c.QueryParam("age")
    return c.String(http.StatusOK, "name:" + name + ", age:" + age)
}

func getUser(c echo.Context) error {
    fmt.Println("getUserが呼ばれました！")
    userId := c.Param("userId")
    return c.String(http.StatusOK, "userId:" + userId)
}

func createUser(c echo.Context) error {
    fmt.Println("createUserが呼ばれました！")
    u := new(User)
    if err := c.Bind(u); err != nil {
       return err
    }
    return c.JSON(http.StatusOK, u)
}
