package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func swap(x, y string) (string, string) {
	return y, x
}

func osFind() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("meow")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("chal bhak 😔")

	}
}

func asign() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

var meo, neow bool

func main() {
	a, b := swap("mr", "fg")
	fmt.Println(a, b)
	defer fmt.Println("nahi tu bhak 😡")
	fmt.Println("chal bhak 😔")
}