package pow

/*
50. Pow(x, n)
实现 pow(x, n) ，即计算 x 的 n 次幂函数。
说明:
-100.0 < x < 100.0
n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
*/

//分治法，要注意边界条件
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n%2 == 1 { //奇数
		return x * myPow(x, n-1)
	}
	return myPow(x*x, n/2) // 偶数
}

//位运算
func myPowBit(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var pow float64 = 1
	for n > 0 {
		if n&1 == 1 { //x奇数
			pow *= x
		}
		x *= x
		n >>= 1
	}
	return pow
}
