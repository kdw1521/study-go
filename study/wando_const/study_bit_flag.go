package wandoconst

import "fmt"

const (
	MasterRoom uint8 = 1 << iota
	BathRoom
	Kitchen
	SmallRoom
)

func SetLight(rooms, room uint8) uint8 {
	return rooms | room
}

func ResetLight(rooms, room uint8) uint8 {
	return rooms &^ room
}

func IsTurnon(rooms, room uint8) bool {
	return rooms&room == room
}

func TurnonLights(rooms uint8) {
	if IsTurnon(rooms, MasterRoom) {
		fmt.Println("마스터룸 불을 킨다!")
	}
	if IsTurnon(rooms, BathRoom) {
		fmt.Println("화장실 불을 킨다!")
	}
	if IsTurnon(rooms, Kitchen) {
		fmt.Println("주방 불을 킨다!")
	}
	if IsTurnon(rooms, SmallRoom) {
		fmt.Println("작은방 불을 킨다!")
	}
}
