package DB

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type WebData struct {
	gorm.Model
	TimeGenerated time.Time `json:"time_generated"`
	Status string `json:"status"`
	Priority int `json:"priority"`
	PayLoad string `json:"pay_load"`
	Country string `json:"country"`
	ClosestDate string `json:"closest_date"`
}
func DbSqlMigration(url string) *gorm.DB {
	db, err := gorm.Open("mysql", url)
	if err != nil {
		log.Println(err)
	}
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 auto_increment=1")
	db.AutoMigrate(&WebData{})
	return db
}
