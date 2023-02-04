package comparison

import "fmt"

func Float_Error() {
	var (
		a float64 = 0.1
		b float64 = 0.2
		c float64 = 0.3
	)

	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)
	fmt.Println(a + b)
}
