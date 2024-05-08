package main

import "fmt"

func main() {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	//fmt.Println(x, y) // compile error: x and y are not visible here

}
func f() int {
	return 1
}
func g(x int) int {
	return 1 + x
}
