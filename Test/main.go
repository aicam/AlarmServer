package main

import (
	"log"
	"strconv"
)

func main() {
	st := "\u0014"
	sti, _ := strconv.Atoi(st)
	log.Print(string(sti))
}
