package thinkMath

import "math"

func Round(f float64, i int) float64 {
	//f += math.Pow10(-15)
	n := math.Pow10(i)
	return math.Trunc((f+0.5/n)*n) / n
}
