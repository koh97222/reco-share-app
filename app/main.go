package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
)

// Sample samplesテーブルに対応する構造体
type Sample struct {
	gorm.Model
	Name sql.NullString
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "docker"
	PASS := "docker"
	PROTOCOL := "@tcp(db:3306)/"
	DBNAME := "recoshare_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return db
}

func main() {
	e := echo.New()
	// e.File("/", "../reco-share/src/App.vue")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})
	e.GET("/sample", func(c echo.Context) error {
		sample := []Sample{}
		db := gormConnect()
		db.Find(&sample)
		fmt.Println(&sample)
		return c.JSON(http.StatusOK, &sample)
	})
	e.Logger.Fatal(e.Start(":8082"))
}
