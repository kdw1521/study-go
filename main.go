package main

import "study-go/study/wando_interface/minju"

// 이제 여기서 wando 말고 minju로 보내고 싶으면 코드를 바꿔야한다..!
func SendBook(name string, sender *minju.MinjuSender) {
	sender.Send(name)
}

func main() {
	sender := &minju.MinjuSender{}
	SendBook("완도 일대기", sender)
	SendBook("완도 일대기2", sender)
}
