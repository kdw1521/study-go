package main

import (
	"study-go/study/solid"
)

func main() {
	var mail = &solid.Mail{}
	var listener solid.EventListener

	listener = &solid.Alarm{}

	mail.Register(listener)
	mail.OnRecv()
}
