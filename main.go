package main

import (
	"container/heap"
	"fmt"
	"strconv"
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
func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var path []int
	var backTrace func(k, n, start, sum int)
	backTrace = func(k, n, start, sum int) {
		if sum == n && len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := start; i <= 9; i++ {
			path = append(path, i)
			sum += i
			backTrace(k, n, i+1, sum)
			sum -= i

			path = path[0 : len(path)-1]
		}
	}
	backTrace(k, n, 1, 0)
	return res
}
func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
func partition(s string) [][]string {
	res := [][]string{}
	path := []string{}
	var backTrace func(sub string)
	backTrace = func(sub string) {
		if sub == "" {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := 1; i <= len(sub); i++ {
			substr := sub[:i]
			if !isPalindrome(substr) {
				return
			}
			path = append(path, substr)
			backTrace(sub[i:])
			path = path[:len(path)-1]
		}
	}
	backTrace(s)
	return res

}

type Person struct {
	name string
	age  int
}

func isIP(ip string) bool {
	if len(ip) > 1 && ip[0:1] == "0" {
		return false
	}
	val, _ := strconv.Atoi(ip)
	if val > 255 {
		return false
	}
	return true
}

func restoreIpAddresses(s string) []string {
	var res []string
	path := ""
	var backTrace func(sub string, count int)
	backTrace = func(sub string, count int) {
		if sub == "" && count == 4 {
			path = path[:len(path)-1]
			res = append(res, path)
			return
		}
		if count > 4 {
			return
		}
		for i := 1; i <= len(sub) && i <= 3; i++ {
			subStr := sub[:i]
			if !isIP(subStr) {
				continue
			}
			count++
			temp := path
			path = path + subStr + "."
			backTrace(sub[i:], count)
			path = temp
			count--
		}
	}
	backTrace(s, 0)
	return res
}

func subsets(nums []int) [][]int {
	var res [][]int
	var path []int
	var backTrace func(startIndex int)
	backTrace = func(startIndex int) {
		res = append(res, path)
		if startIndex >= len(nums) {
			return
		}
		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			backTrace(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTrace(0)
	return res
}

func solveNQueens(n int) [][]string {
	var res [][]string
	path := make([]string, n)
	tempRow := ""
	for i := 0; i < n; i++ {
		tempRow += "."
	}
	for i := 0; i < n; i++ {
		path[i] = tempRow
	}
	var backTrace func(row int)
	backTrace = func(row int) {
		if row == n {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		for col := 0; col < n; col++ {
			if isVaild(row, col, n, path) {
				path[row] = path[row][:col] + "Q" + path[row][col+1:]
				backTrace(row + 1)
				path[row] = tempRow
			}
		}
	}
	backTrace(0)
	return res
}

func isVaild(row, col, n int, path []string) bool {
	for i := 0; i < row; i++ {
		if path[i][col:col+1] == "Q" {
			return false
		}
	}
	for i, j := col-1, row-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if path[j][i:i+1] == "Q" {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if path[i][j:j+1] == "Q" {
			return false
		}
	}
	return true
}
func isTVaild(board [][]byte, row, col int, n byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == n || board[i][col] == n {
			return false
		}
	}
	startRow := row / 3 * 3
	startCol := col / 3 * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == n {
				return false
			}
		}
	}
	return true
}

func solveSudoku(board [][]byte) {
	var backTrace func() bool
	backTrace = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' {
					continue
				}
				var k byte
				for k = '1'; k <= '9'; k++ {
					if isTVaild(board, i, j, k) {
						board[i][j] = k
						if backTrace() {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}
	backTrace()
}
func main() {
	subsets([]int{1, 2, 3})
	fmt.Println(solveNQueens(4))
}
