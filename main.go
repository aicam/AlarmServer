package AlarmServer

import (
	"github.com/aicam/AlarmServer/DB"
)

func main() {
	// migration
	s := NewServer()
	s.DB = DB.DbSqlMigration("aicam:021021ali@tcp(127.0.0.1:3306)/messenger_api?charset=utf8mb4&parseTime=True")

}
