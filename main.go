package main

import (
	"github.com/aicam/AlarmServer/DB"
	"github.com/aicam/AlarmServer/server"
	"log"
	"net/http"
)

func main() {
	// migration
	s := server.NewServer()
	s.DB = DB.DbSqlMigration("aicam:021021ali@tcp(127.0.0.1:3306)/messenger_api?charset=utf8mb4&parseTime=True")
	s.Routes()
	err := http.ListenAndServe("0.0.0.0:4300", s.Router)
	if err != nil {
		log.Print(err)
	}

}
