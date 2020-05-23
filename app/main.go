package main

import (
	"app/db"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// User ログイン時に送信される情報に対応する構造体
type User struct {
	UserID int    `gorm:"column:user_id"`
	User   string `gorm:"column:user_nm"`
	Pass   string `gorm:"column:password"`
	Email  string `gorm:"column:email"`
}

// Response APIのレスポンス結果に対応する構造体
type Response struct {
	ResultCode    string `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}

func guestLogin() User {
	db := db.GormConnect()
	user := User{}
	db.Where("user_id = ?", "1").Find(&user)
	defer db.Close()
	return user
}

// registUser　ユーザ登録メソッド
func registUser(u User) {
	db := db.GormConnect()
	db.Create(u)
	defer db.Close()
}

// userDuplicateCheck ユーザ名重複確認メソッド
func userDuplicateCheck(u User) bool {
	db := db.GormConnect()
	user := User{}
	db.Where("user_nm = ?", u.User).Find(&user)
	defer db.Close()
	if user.UserID == 0 {
		return false
	}
	return true
}

// registUserHandler ユーザ登録のハンドラ
func registUserHandler(c echo.Context) error {
	r := Response{}
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	// 既に同じユーザ名が登録されていない場合、登録を行う。
	if !userDuplicateCheck(*user) {
		registUser(*user)
		r.ResultCode = "00"
		r.ResultMessage = "registerd"
	} else {
		r.ResultCode = "80"
		r.ResultMessage = "Not registerd"
	}
	return c.JSON(http.StatusOK, r)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
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
	e.POST("/registuser", registUserHandler)
	e.Logger.Fatal(e.Start(":8082"))
}
