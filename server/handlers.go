package server

import (
	"github.com/aicam/AlarmServer/DB"
	"github.com/gin-gonic/gin"
	"net/http"
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
		if timeFounded.Sub(time.Now()).Hours()/24 < 10 {
			go sendNotificationByPushOver("Time found in "+timeFounded.Month().String()+" "+string(timeFounded.Day()), "Time found less than "+
				string(int(timeFounded.Sub(time.Now()).Hours()/24))+" days"+" in "+jsData.Country)
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
