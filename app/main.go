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
	UserNm string `gorm:"column:user_nm"`
	Pass   string `gorm:"column:password"`
	Email  string `gorm:"column:email"`
}

// Response APIのレスポンス結果に対応する構造体
type Response struct {
	ResultCode    string `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
	User
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

// isUserAlreadyRegistered ユーザ名重複確認メソッド
func isUserAlreadyRegistered(u User) bool {
	db := db.GormConnect()
	user := User{}
	recordNotFound := db.Where("user_nm = ?", u.UserNm).Find(&user).RecordNotFound()
	defer db.Close()
	if recordNotFound {
		return false
	}
	return true
}

// login ログインメソッド
// ここでは簡易的にユーザ名とパスワードが合っている場合、trueを返す。
// TODO:認証機能を実装すること。
func login(u User) bool {
	db := db.GormConnect()
	user := User{}
	recordNotFound := db.Where("user_nm = ? AND password = ?", u.UserNm, u.Pass).Find(&user).RecordNotFound()
	defer db.Close()
	if recordNotFound {
		return false
	}
	return true
}

// getUserInfo ユーザ情報取得
// ユーザ名に紐づく情報を取得し、返却する。
func getUserInfo(userNm string) User {
	db := db.GormConnect()
	user := User{}
	db.Where("user_nm = ?", userNm).Find(&user)
	defer db.Close()
	return user
}

// registUserHandler ユーザ登録のハンドラ
func registUserHandler(c echo.Context) error {
	r := Response{}
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	// 既に同じユーザ名が登録されていない場合、登録を行う。
	if !isUserAlreadyRegistered(*user) {
		registUser(*user)
		r.ResultCode = "00"
		r.ResultMessage = "registerd"
	} else {
		r.ResultCode = "80"
		r.ResultMessage = "Not registerd"
	}
	return c.JSON(http.StatusOK, r)
}

// loginHandler ログインのハンドラ
func loginHandler(c echo.Context) error {
	r := new(Response)
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	// ログイン成功時、ユーザ情報をさらに取得し、返却をする。
	if login(*user) {
		user := getUserInfo(user.UserNm)
		r.ResultCode = "00"
		r.ResultMessage = "login success"
		r.UserID = user.UserID
		r.UserNm = user.UserNm
	} else {
		// ログイン失敗時
		r.ResultCode = "80"
		r.ResultMessage = "login failure"
	}
	return c.JSON(http.StatusOK, r)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})
	e.POST("/login", loginHandler)
	e.POST("/guestlogin", func(c echo.Context) error {
		return c.JSON(http.StatusOK, guestLogin())
	})
	e.POST("/registuser", registUserHandler)
	e.Logger.Fatal(e.Start(":8082"))
}
