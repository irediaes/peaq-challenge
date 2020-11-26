package utils

// CalculatePercentageDifference ...
func CalculatePercentageDifference(num1, num2 float64) float64 {

	diff := difference(num1, num2)
	average := (num1 + num2) / 2

	percentDiff := (diff / average) * 100

	return percentDiff
}

func difference(num1, num2 float64) float64 {
	if num1 > num2 {
		return (num1 - num2)
	}
	return (num2 - num1)
}
