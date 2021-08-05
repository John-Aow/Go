package main

import (
	"fmt"
	"math"
)

var index int
var task map[int]string = map[int]string{}

func main() {
	fmt.Println(Power(2, 3))
}

func Power(base, exponent int) int {
	return int(math.Pow(float64(base), float64(exponent)))
}
