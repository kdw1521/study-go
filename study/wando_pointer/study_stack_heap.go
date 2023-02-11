package wandopointer

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
} // 여기서 u 는 사라지는데 그렇다면 사라진 메모리주소를 반환한다.

func Runner() {
	userPointer := NewUser("wando", 31)
	fmt.Println(userPointer) // 그럼 이게 어떻게 나오는거지?
}

/*
일단 func 내에 있는 지역변수는 stack에 쌓인다.
함수가 끝날때 stack에서 pop을 해버린다.

하지만, go 에서는 Escape Analysis(탈출분석) 을 한다.
userPointer := NewUser("wando", 31) 이와 같이 NewUser에서 만들어진 인스턴스가 바깥으로 탈출하는 형태로 볼수있다.
그래서 이건 stack에 할당 되면 안된다.
그래서 고 컴파일러가 분석을 해서 힙(Heap)에 할당해둔다.
힙은 쓰임이 없어질때 GC가 지워준다.
*/
