package DB

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type WebData struct {
	gorm.Model
	TimeGenerated string `json:"time_generated"`
	Status        string `json:"status"`
	Priority      int    `json:"priority"`
	PayLoad       string `json:"pay_load"`
	Country       string `json:"country"`
	ClosestDate   string `json:"closest_date"`
}

func DbSqlMigration(url string) *gorm.DB {
	db, err := gorm.Open("mysql", url)
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&WebData{})
	return db
}
