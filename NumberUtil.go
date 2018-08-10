package util

import (
	"math"
	"strconv"
)

func KeepDecimal(num float64, i int) float64 {
	n := math.Pow10(i)
	x := int(num*n + 0.5)
	// 四舍五入
	return float64(x) / n
}

func FormatFloat(f float64, decimal int) string {
	return strconv.FormatFloat(f, 'g', decimal+1, 64)
}
