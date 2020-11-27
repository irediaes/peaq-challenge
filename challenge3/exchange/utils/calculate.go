package utils

import (
	"fmt"
	"strconv"
)

// CalculatePercentageDifference ...
func CalculatePercentageDifference(newNum, originNum float64) float64 {

	diff := newNum - originNum

	percentDiff := (diff / originNum) * 100

	return Round4Decimal(percentDiff)
}

// Round4Decimal ...
func Round4Decimal(num float64) float64 {
	numb := fmt.Sprintf("%.4f", num)
	result, _ := strconv.ParseFloat(numb, 64)
	if result == 0 {
		return 0.0
	}
	return result
}
