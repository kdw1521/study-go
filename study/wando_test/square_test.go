package wandotest

import "testing"

func TestSquare(t *testing.T) {
	result := square(9)
	if result != 81 {
		t.Errorf("기대깂 :81, 출력값:%d\n", result)
	}
}

func TestSquare2(t *testing.T) {
	result := square(3)
	if result != 9 {
		t.Errorf("기대깂 :9, 출력값:%d\n", result)
	}
}
