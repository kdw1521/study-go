package wandointerface

import "fmt"

type Attacker interface {
	Attack()
}

type Monster struct {
	Lv int
}

func (m *Monster) Attack() {
	fmt.Println("Monster Attack")
}

func DoAttack(att Attacker) {
	if att != nil {
		att.Attack()

		var monster *Monster
		monster = att.(*Monster)
		fmt.Println(monster.Lv)

	}
}
