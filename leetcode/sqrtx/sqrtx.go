package sqrtx

/* 69. x 的平方根
实现 int sqrt(int x) 函数。
计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

链接：https://leetcode-cn.com/problems/sqrtx
*/
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right := 0, x
	var res int
	for left <= right {
		m := (left + right) / 2
		if m == x/m {
			return m
		} else if m > x/m {
			right = m - 1
		} else {
			left = m + 1
			res = m
		}
	}
	return res
}
