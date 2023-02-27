package wandointerface

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func (s Student) GetAge() int { // 요놈은 Stringer interface에서는 당연히 호출이 안됨!
	return s.Age
}

func Runner() {
	student := Student{"wando", 31}
	var stringer Stringer

	stringer = student
	fmt.Printf("%s\n", stringer.String())
}
