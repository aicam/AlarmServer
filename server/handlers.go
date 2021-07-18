package server

import (
	"github.com/aicam/AlarmServer/DB"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
		log.Print(jsData)
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
