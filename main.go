package main

import (
	"fmt"
)

var result [][]int
var i int = 1

func main() {
	res := solveNQueens(5)
	fmt.Println(res)

}

func solveNQueens(n int) [][]string {
	if n == 1 {
		return [][]string{{"Q"}}
	}
	if n <= 3 {
		return nil
	}
	// result = nil
	var cols map[int]byte
	var pie map[int]byte
	var na map[int]byte
	var cur_state []int
	DFS(n, 0, cols, pie, na, cur_state)
	return generate_result(n)
}

func DFS(n int, row int, cols map[int]byte, pie map[int]byte, na map[int]byte, cur_state []int) {
	//recursion terminator
	if row >= n {
		fmt.Println("i:", i, " state:", cur_state)
		result = append(result, cur_state)
		i++
	}
	for col := 0; col < n; col++ {
		if row == 0 {
			fmt.Println("init ...")
			cols = make(map[int]byte)
			pie = make(map[int]byte)
			na = make(map[int]byte)
			cur_state = nil
		}
		_, ok1 := cols[col]
		_, ok2 := pie[row+col]
		_, ok3 := na[row-col]
		if ok1 || ok2 || ok3 {
			continue
		}

		//get the postion,update the flags
		cols[col] = '1'
		pie[row+col] = '1'
		na[row-col] = '1'

		cur_state = append(cur_state, col)
		//go to next row
		DFS(n, row+1, cols, pie, na, cur_state)
	}
}
func generate_result(n int) [][]string {
	var res [][]string
	for i := 0; i < len(result); i++ {
		r := make([]string, n)
		for j := 0; j < n; j++ {
			r[j] = _generate(n, result[i][j])

		}
		res = append(res, r)
	}
	return res
}

func _generate(n int, col int) string {
	var pos []byte
	for i := 0; i < col; i++ {
		pos = append(pos, '.')
	}
	pos = append(pos, 'Q')
	for j := 0; j < n-col-1; j++ {
		pos = append(pos, '.')
	}
	return string(pos)
}
