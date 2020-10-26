package queens

import "strings"

// import mapset "github.com/deckarep/golang-set"

/*51. N 皇后
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。
链接：https://leetcode-cn.com/problems/n-queens
示例：
输入：4
输出：[
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]
解释: 4 皇后问题存在两个不同的解法。
*/
var ans [][]string

var cols map[int]bool
var pie map[int]bool
var na map[int]bool

func solveNQueens(n int) [][]string {
	if n <= 0 {
		return ans
	}
	ans = [][]string{}

	cols = make(map[int]bool, n)
	pie = make(map[int]bool, n)
	na = make(map[int]bool, n)

	dfs(n, 0, []int{})
	return ans
}

func dfs(n int, row int, path []int) {
	if row >= n {
		ans = append(ans, generateBoard(path, n))
		return
	}

	for col := 0; col < n; col++ {
		if cols[col] || pie[row+col] || na[row-col] {
			continue
		}
		cols[col] = true
		pie[row+col] = true
		na[row-col] = true
		dfs(n, row+1, append(path, col))
		// cols pie na 都是全局变量，所以需要恢复现场供下次循环使用
		cols[col] = false
		pie[row+col] = false
		na[row-col] = false
	}
	return
}

func generateBoard(path []int, n int) []string {
	var res []string
	for _, col := range path {
		res = append(res, strings.Repeat(".", col)+"Q"+strings.Repeat(".", n-col-1))
	}
	return res
}
