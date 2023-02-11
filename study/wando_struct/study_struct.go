package wandostruct

import "fmt"

type House struct {
	Address  string
	Size     int
	Price    float64
	Catefory string
}

func Basic_Struct() {
	var house House
	house.Address = "서울 강남구 ..."
	house.Size = 33
	house.Price = 10
	house.Catefory = "아파트"

	fmt.Println(house)
}

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func Innner_Struct() {
	user := User{"완도", "wando", 31}
	vip := VIPUser{
		User{"로만두", "min", 31},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP유저: %s ID: %s 나이: %d VIP레벨: %d VIP가격: %d만원\n", vip.UserInfo.Name, vip.UserInfo.ID, vip.UserInfo.Age, vip.VIPLevel, vip.Price)
}

type EmbeddedUser struct {
	Name string
	ID   string
	Age  int
}

type EmbeddedVIPUser struct {
	EmbeddedUser // 내장된 필드
	VIPLevel     int
	Price        int
	Name         string
}

func Embedded_Struct() {
	vip := EmbeddedVIPUser{
		EmbeddedUser{"로만두", "min", 31},
		3,
		250,
		"wando",
	}

	fmt.Printf("VIP유저: %s ID: %s 나이: %d VIP레벨: %d VIP가격: %d만원\n", vip.Name, vip.ID, vip.Age, vip.VIPLevel, vip.Price)
}
