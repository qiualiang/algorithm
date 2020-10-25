package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(math.Dim(0.25, 0.25000) < 0.00001)
	fmt.Println(math.Dim(9.261000000000001, 9.26100) < 0.00001)

	var ff float64
	ff = 0.2500
	ff = FloatRound(ff, 5)
	fmt.Printf("%f", ff) // 输出 -1.3551
}

// 截取小数位数
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}
