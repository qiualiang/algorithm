package parentheses

/*22. 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]

链接：https://leetcode-cn.com/problems/generate-parentheses
*/
var list []string

func generateParenthesis(n int) []string {
	list = nil
	gen(0, 0, n, "")
	return list
}

//DFS
func gen(left int, right int, n int, result string) {
	//terminate
	if left == n && right == n {
		list = append(list, result)
		return
	}
	if left < n {
		gen(left+1, right, n, result+"(")
	}
	if right < n && left > right {
		gen(left, right+1, n, result+")")
	}
}
