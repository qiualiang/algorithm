package main

import (
	"fmt"
)

func main() {
	s := []int{0}
	s1 := s[:len(s)-1]
	fmt.Println(s1)
}
