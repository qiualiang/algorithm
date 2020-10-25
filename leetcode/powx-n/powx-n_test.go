package pow

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestPow1(t *testing.T) {
	result := myPow(2.00000, 10)
	var p float64 = 0.00001
	if !isEqual(result, 1024.00000, p) {
		t.Error("myPow(2.0000, 10) returns:", result)
	}

	result = myPow(2.10000, 3)
	if !isEqual(result, 9.26100, p) {
		t.Error("myPow(2.10000, 3) returns:", result)
	}
	result = myPow(2.00000, -2)
	if !isEqual(result, 0.25000, p) {
		t.Error("myPow(2.00000, -2) returns:", result)
	}
}
func TestPowBit(t *testing.T) {
	result := myPowBit(2.00000, 10)
	var p float64 = 0.00001
	if !isEqual(result, 1024.00000, p) {
		t.Error("myPow(2.0000, 10) returns:", result)
	}

	result = myPowBit(2.10000, 3)
	if !isEqual(result, 9.26100, p) {
		t.Error("myPow(2.10000, 3) returns:", result)
	}
	result = myPowBit(2.00000, -2)
	if !isEqual(result, 0.25000, p) {
		t.Error("myPow(2.00000, -2) returns:", result)
	}
}

//p为自定精度:0.00001
func isEqual(f1, f2, p float64) bool {

	return math.Dim(f1, f2) < p
}

// 截取小数位数
func floatRound(f float64, precision int) float64 {
	format := "%." + strconv.Itoa(precision) + "f"
	fmt.Println(format)
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}
