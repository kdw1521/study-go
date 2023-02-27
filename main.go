package main

import (
	// "fmt"
	"study-go/study/wando_interface"
	// "unsafe"
)

// type Sender interface {
// 	Send(parcel string)
// }

// func SendBook(name string, sender Sender) {
// 	sender.Send(name)
// }

func main() {
	// sender := &wandointerface.WandoSender{} // or &minju.MinjuSender{}
	// SendBook("완도 일대기", sender)
	// SendBook("완도 일대기2", sender)

	// var stringer2_1 wandointerface.Stringer2
	// fmt.Printf("type:%T size:%d\n", stringer2_1, unsafe.Sizeof(stringer2_1))

	// student := &wandointerface.Student2{Name: "WANDO"}
	// stringer2_1 = student
	// fmt.Printf("type:%T size:%d\n", stringer2_1, unsafe.Sizeof(stringer2_1))

	// var stringer2_2 wandointerface.Stringer2
	// fmt.Printf("type:%T size:%d\n", stringer2_2, unsafe.Sizeof(stringer2_2))

	// user := wandointerface.User{Name: "WANDO@", Age: 31}
	// stringer2_2 = user
	// fmt.Printf("type:%T size:%d\n", stringer2_2, unsafe.Sizeof(stringer2_2))

	// wandointerface.PrintVal(10)
	// wandointerface.PrintVal(3.14)
	// wandointerface.PrintVal("Wando")
	// wandointerface.PrintVal(wandointerface.People{31})

	// var att wandointerface.Attacker
	// att.Attack()
	// wandointerface.DoAttack(&wandointerface.Monster{Lv: 20})

	// var wando wandointerface.Wando
	// wando.(*wandointerface.DoDo)

	file := &wandointerface.File{}
	wandointerface.ReadFile(file)
}
