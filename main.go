package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"worms/utils"
)

type Point utils.Point

func test(a *[]int) []int {
	b := *a
	b[1] = 2
	*a = append(*a, 4)
	return *a
}

func (p Point) name() {
	fmt.Println(p.X, p.Y)
}
func topKFrequent(nums []int, k int) []int {
	map_num := map[int]int{}
	//记录每个元素出现的次数
	for _, item := range nums {
		map_num[item]++
	}
	h := &IHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	for key, value := range map_num {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	//按顺序返回堆中的元素
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}

//构建小顶堆
type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	point := new(utils.Point)
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
	point.X = 6
	point.Y = 9
	point.GetDistance(*new(utils.Point), *new(utils.Point))
	var k = *new(Point)
	k.X = 9
	k.Y = 10
	k.name()
	fmt.Println(c, d, f, q[0:2], &s, &p, len(*p), &d, nil, *point)
	a := []int{1, 2}
	test(&a)
	fmt.Println(a)
	l := list.New()
	l.Back()

}
