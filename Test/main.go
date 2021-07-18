package main

import (
	"log"
	"time"
)

func main() {
	location, err := time.LoadLocation("Local")
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(time.Date(2021, time.Month(2), 3, 0, 0, 0, 0, location))
}
