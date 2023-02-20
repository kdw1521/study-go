package main

import (
	"study-go/study/wando_interface"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	sender := &wandointerface.WandoSender{} // or &minju.MinjuSender{}
	SendBook("완도 일대기", sender)
	SendBook("완도 일대기2", sender)
}
