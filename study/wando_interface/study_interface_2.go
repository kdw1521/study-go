package wandointerface

type Stringer2 interface {
	String2() string
}

type Student2 struct {
	Name string
}

func (s *Student2) String2() string {
	return s.Name
}

type User struct {
	Name string
	Age  int
}

func (u User) String2() string {
	return u.Name
}
