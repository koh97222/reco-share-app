package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/koh97222/reco-share-app/app/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Sample samplesテーブルに対応する構造体
type Sample struct {
	Name string `gorm:"column:name"`
	gorm.Model
}

func getSample() []Sample {
	db := db.GormConnect()
	sample := []Sample{}
	db.Find(&sample)
	defer db.Close()
	return sample
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
	e.Logger.Fatal(e.Start(":8082"))
}
