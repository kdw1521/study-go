package wandointerface

import "fmt"

type WandoSender struct {
}

// 외부에서 이걸 이용해서 택배를 보낸다.
func (w *WandoSender) Send(parcel string) {
	fmt.Printf("Wando sends %s parcel\n", parcel)
}
