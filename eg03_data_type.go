package main

import (
	"fmt"
)

func main() {
	// var name string = "Tom"
	// age := 20
	// m := true
	// fmt.Printf("%T\n", name)
	// fmt.Printf("%T\n", age)
	// fmt.Printf("%T\n", m)
	a := 10
	p := &a
	fmt.Printf("%T\n", p)

	arr := [2]int{1, 2}
	fmt.Printf("%T\n", arr)

	arr_1 := []int{1, 2, 3}
	fmt.Printf("%T\n", arr_1)

}
