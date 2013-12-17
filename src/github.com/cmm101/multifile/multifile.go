package multifile

import (
	"fmt"
	"math"
)

func Helloworld() {
	fmt.Println("\nHello world From File 2.")
	fmt.Println(math.Abs(30))
	fmt.Println(add(42, 13))
}
func add(x int, y int) int {
	return x + y
}
