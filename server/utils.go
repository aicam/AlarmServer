package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func convertMonthToInt(month string) int {
	var m map[string]int
	m = make(map[string]int)
	m["January"] = 1
	m["February"] = 2
	m["March"] = 3
	m["April"] = 4
	m["May"] = 5
	m["June"] = 6
	m["July"] = 7
	m["August"] = 8
	m["September"] = 9
	m["October"] = 10
	m["November"] = 11
	m["December"] = 12
	return m[month]
}

type NotifReq struct {
	Token   string `json:"token"`
	User    string `json:"user"`
	Message string `json:"message"`
	Title   string `json:"title"`
}

func sendNotificationByPushOver(message string, title string) {
	url := "https://api.pushover.net/1/messages.json"

	var jsonBytes []byte
	jsonBytes, err := json.Marshal(&NotifReq{
		Token:   "",
		User:    "",
		Message: message,
		Title:   title,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonBytes))
	//req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println(err)
		return
	}
	//client := &http.Client{}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
