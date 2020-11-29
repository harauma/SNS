package main

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

const (
    // データベース
    Dialect = "mysql"

    // ユーザー名
    DBUser = "sisupiyo"

    // パスワード
    DBPass = "sisupiyo"

    // プロトコル
    DBProtocol = "tcp(mysql:3306)"

    // DB名
    DBName = "test_db"
)

//DBのuserテーブルの構造に合わせる
type User struct {
    Id string `json:"id"`
    Name  string `json:"name"`
}

func main() {
    e := echo.New()
    initRouting(e)
    e.Logger.Fatal(e.Start(":8000"))
}

//ルーティングの作成
func initRouting(e *echo.Echo) {
    e.GET("/", echoHello)
    e.GET("/users", getUsers)
    e.GET("/user/:userId", getUser)
    e.POST("/user", createUser)
    // e.PUT("/user", updateUser)
    e.DELETE("/user/:userId", deleteUser)
}

//DB接続
func connectGorm() *gorm.DB {
    connectTemplate := "%s:%s@%s/%s"
    connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
    db, err := gorm.Open(Dialect, connect)

    if err != nil {
        fmt.Println(err.Error())
    }

    return db
}

func echoHello(c echo.Context) error {
    fmt.Println("echoHelloが呼ばれました！")
    return c.String(http.StatusOK, "Hello motio")
}

func getUsers(c echo.Context) error {
    fmt.Println("getUsersが呼ばれました！")

    db := connectGorm()
    db.SingularTable(true)
    defer db.Close()

    result := findAll(db)
    fmt.Println(result)

    // name := c.QueryParam("name")
    // age := c.QueryParam("age")

    // return c.String(http.StatusOK, "name:" + name + ", age:" + age)
    return c.JSON(http.StatusOK, result)
}

func getUser(c echo.Context) error {
    fmt.Println("getUserが呼ばれました！")
    db := connectGorm()
    db.SingularTable(true)
    defer db.Close()

    userId := c.Param("userId")
    result := findUser(db, userId)

    // return c.String(http.StatusOK, "userId:" + userId)
    return c.JSON(http.StatusOK, result)
}

//user新規作成
func createUser(c echo.Context) error {
    fmt.Println("createUserが呼ばれました！")
    db := connectGorm()
    db.SingularTable(true)
    defer db.Close()

    u := new(User)
    if err := c.Bind(u); err != nil {
        return err
    }
    db.Create(&u)

    return c.JSON(http.StatusOK, u)
}

//user削除
func deleteUser(c echo.Context) error {
    fmt.Println("createUserが呼ばれました！")
    db := connectGorm()
    db.SingularTable(true)
    defer db.Close()

    userId := c.Param("userId")
    u := new(User)
    if err := c.Bind(u); err != nil {
        return err
    }

    if err := db.Where("id = ?", userId).Delete(&u).Error; err != nil {
        fmt.Println(err)
        return (c.JSON(http.StatusNotFound, nil))
    }

    return c.String(http.StatusNoContent, "id : " + userId + " deleted")
}

//userテーブルの全件検索
func findAll(db *gorm.DB) []User {
    var allUsers []User
    db.Find(&allUsers)
    return allUsers
}

//userテーブルからidを指定して検索
func findUser(db *gorm.DB, id string) User {
    var user User
    db.Where("Id = ?",id).Find(&user)
    return user
}
