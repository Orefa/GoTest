package main

import (
	"fmt"
	"unsafe"
)

type PPoint struct {
	x, y int
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		fmt.Print(a[i])
	}
	fmt.Println()
	s := a[1:3]
	s[0] = 0
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i])
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Print(a[i])
	}
	fmt.Println()
	ss := append(s, 6, 7, 8)
	for i := 0; i < len(ss); i++ {
		fmt.Print(ss[i])
	}
	fmt.Println()
	ss[0] = 9

	for i := 0; i < len(ss); i++ {
		fmt.Print(ss[i])
	}
	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i])
	}
	p := make([]byte, 2, 6)
	fmt.Println()
	fmt.Println("p sizeOf=", unsafe.Sizeof(&p))
	pp := new(PPoint)

	fmt.Println(unsafe.Sizeof(*pp))
}
