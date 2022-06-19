package main

import (
	"fmt"
	"worms/utils"
)

func main() {
	fmt.Println("main")
	fmt.Println(utils.Add(4, 5))
	var c string = "test"
	var d int = 9
	var f int8 = 10
	f = int8(d)
	c = c[1:3]
	q := []int8{1, 23}
	q[1] = 3
	s := make([]int16, 2, 4)
	p := new([]int16)
	*p = append(*p, 3)
	fmt.Println(c, d, f, q[0:2], &s, &p, len(*p))
}
