package main

import (
	"app/db"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Sample samplesテーブルに対応する構造体
type Sample struct {
	Name string `gorm:"column:name"`
	gorm.Model
}

// User ログイン時に送信される情報に対応する構造体
type User struct {
	UserID int    `gorm:"column:user_id"`
	User   string `gorm:"column:user_nm"`
	Pass   string `gorm:"column:password"`
	Email  string `gorm:"column:email"`
}

func getSample() []Sample {
	db := db.GormConnect()
	sample := []Sample{}
	db.Find(&sample)
	defer db.Close()
	return sample
}

func guestLogin() User {
	db := db.GormConnect()
	user := User{}
	db.Where("user_id = ?", "1").Find(&user)
	defer db.Close()
	return user
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})
	e.GET("/sample", func(c echo.Context) error {
		return c.JSON(http.StatusOK, getSample())
	})
	// e.POST("/login", func(c echo.Context) error {
	// 	login := User{}
	// 	login.User = c.QueryParam("users")
	// 	login.Pass = c.QueryParam("users.pass")
	// 	return c.JSON(http.StatusOK, login)
	// })
	e.POST("/guestlogin", func(c echo.Context) error {
		return c.JSON(http.StatusOK, guestLogin())
	})
	e.Logger.Fatal(e.Start(":8082"))
}
