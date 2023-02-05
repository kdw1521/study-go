package main

import (
	wandoconst "study-go/study/wando_const"
)

func main() {
	var rooms uint8
	rooms = wandoconst.SetLight(rooms, wandoconst.MasterRoom)
	rooms = wandoconst.SetLight(rooms, wandoconst.SmallRoom)
	rooms = wandoconst.ResetLight(rooms, wandoconst.MasterRoom)
	wandoconst.TurnonLights(rooms)
}
