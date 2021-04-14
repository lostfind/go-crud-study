package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	db = DbConnect()
	defer db.Close()

	RunServer()
}

func DbConnect() *sql.DB {
	db, err := sql.Open("mysql", "writeitgo:password@tcp(db:3306)/go_study")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func RunServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		data := GetData(db)

		c.JSON(200, data)
	})
	r.GET("/all", func(c *gin.Context) {

		data := GetDatas(db)

		c.JSON(200, data)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type AddressData struct {
	ZipCode      string `json:"zip_code"`
	PrefID       int64  `json:"pref_id"`
	PrefName     string `json:"pref_name"`
	PrefNameKana string `json:"pref_name_kana"`
	CityID       int64  `json:"city_id"`
	CityName     string `json:"city_name"`
	CityNameKana string `json:"city_name_kana"`
	TownName     string `json:"town_name"`
	TownNameKana string `json:"town_name_kana"`
}

func GetData(db *sql.DB) AddressData {
	data := AddressData{}
	result := db.QueryRow("SELECT zipcode, prefecture_id, prefecture_name, prefecture_kana, city_id, city_name, city_kana, town_name, town_kana FROM zipcode WHERE zipcode='4630062'")

	result.Scan(&data.ZipCode, &data.PrefID, &data.PrefName, &data.PrefNameKana, &data.CityID, &data.CityName, &data.CityNameKana, &data.TownName, &data.TownNameKana)

	return data
}

func GetDatas(db *sql.DB) []AddressData {
	datas := []AddressData{}
	results, err := db.Query("SELECT zipcode, prefecture_id, prefecture_name, prefecture_kana, city_id, city_name, city_kana, town_name, town_kana FROM zipcode")
	if err != nil {
		log.Fatal(err)
	}

	for results.Next() {
		data := AddressData{}
		results.Scan(&data.ZipCode, &data.PrefID, &data.PrefName, &data.PrefNameKana, &data.CityID, &data.CityName, &data.CityNameKana, &data.TownName, &data.TownNameKana)

		datas = append(datas, data)
	}

	return datas
}
