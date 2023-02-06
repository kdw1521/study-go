package condition

import "fmt"

func Switch_condition() {
	day := "월"

	switch day {
	case "월", "화", "수":
		fmt.Println("월, 화, 수 중에 하나")
	case "목", "금", "토", "일":
		fmt.Println("목,금,토,일 중에 하나")
	}
}

func Switch_If_condition() {
	test := 18

	switch true {
	case test < 10, test > 30:
		fmt.Println("11")
	case test >= 10 && test < 20:
		fmt.Println("22")
	default:
		fmt.Println("default!")
	}
}

type ColorType int

const (
	Violet ColorType = iota
	Red
	Green
	Blue
)

func ColorToString(color ColorType) string {
	switch color {
	case Violet:
		return "Violet"
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	default:
		return "Undifined"
	}
}

func GetMyFavoriteColor() ColorType {
	return Violet
}

func Switch_Fallthrough() {
	a := 3

	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("default")
	}
}
