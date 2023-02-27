package minju

import "fmt"

type MinjuSender struct {
}

func (m *MinjuSender) Send(parcel string) {
	fmt.Printf("민주가 택배 %s 를 보내요.\n", parcel)
}
