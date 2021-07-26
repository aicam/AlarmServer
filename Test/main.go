package main

import "github.com/aicam/AlarmServer/server"

func main() {
	server.SendNotificationByTelegram("Test", "Test PostMan")
}
