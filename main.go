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

type Vertex struct {
	Lat, Long float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(34)
}

var m map[string]Vertex

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

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("alu")
	case string:
		fmt.Printf("is %v tamatar", len(v))
	default:
		fmt.Println("hahaha")
	}
}

func asign() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

var meo, neow bool

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	a, b := swap("mr", "fg")
	fmt.Println(a, b)
	defer fmt.Println("nahi tu bhak 😡")
	fmt.Println("chal bhak 😔")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	primes := [6]int{34, 34, 65, 23, 45, 3}
	var s []int = primes[1:4]
	fmt.Println(s)

	u := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, true},
		{4, true},
		{5, true},
		{6, true},
	}
	fmt.Println(u)

	// make is used for initalizing reference type
	// 0 value of reference are nil

	h := make([]int, 5)
	printSlice("h", h)
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.2837, -82.232,
	}
	fmt.Println(m["Bell Labs"])

	v, ok := m["Labs"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("chal bhak 😔 maps")
	}

	manyEmojies := func() string {
		return "😡😏😎"
	}

	fmt.Println(manyEmojies())
	//NOTE: while using the methods on the struct always remeber that if once you have used the pointer
	//based method do not used the the
	// the byValue method because it is not good to use both at the same time for a single struct

}
