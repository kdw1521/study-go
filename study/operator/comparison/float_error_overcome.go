package comparison

import (
	"fmt"
	"math"
)

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func Overcome_Float_Error() {
	var (
		a float64 = 0.1
		b float64 = 0.2
		c float64 = 0.3
	)

	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))
}
