package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

var meo, neow bool

func main() {
	a, b := swap("mr", "fg")
	fmt.Println(a, b)
}
