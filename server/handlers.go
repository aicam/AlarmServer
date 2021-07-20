package server

import (
	"github.com/aicam/AlarmServer/DB"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	Body       string `json:"body"`
}

func (s *Server) AddInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		var jsData DB.WebData
		err := context.BindJSON(&jsData)
		if err != nil {
			context.JSON(http.StatusOK, Response{
				StatusCode: -1,
				Body:       err.Error(),
			})
			return
		}
		layout := "2006-01-02T15:04:05Z07:00"
		timeFounded, err := time.Parse(layout, jsData.ClosestDate)
		if err != nil {
			context.JSON(http.StatusOK, Response{
				StatusCode: -1,
				Body:       err.Error(),
			})
			return
		}
		if jsData.Priority > 0 {
			log.Print(strconv.Itoa(int(timeFounded.Sub(time.Now()).Hours() / 24)))
			go sendNotificationByPushOver("In "+timeFounded.Month().String()+" "+strconv.Itoa(timeFounded.Day()), "Time found in "+
				strconv.Itoa(int(timeFounded.Sub(time.Now()).Hours()/24))+" days"+" in "+jsData.Country)
		}
		s.DB.Save(jsData)
		context.JSON(http.StatusOK, Response{
			StatusCode: 1,
			Body:       "Data saved successfully!",
		})
	}
}

func (s *Server) GetInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		var DBData []DB.WebData
		s.DB.Find(&DBData)
		context.JSON(http.StatusOK, DBData)
	}
}
