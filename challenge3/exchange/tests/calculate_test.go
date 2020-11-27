package test

import (
	"testing"

	"github.com/ebikode/peaq-challenge/challenge3/exchange/utils"
)

// TestMain ...
func TestCalculate(t *testing.T) {

	originNum := 0.00000857
	newNum := 0.00000740

	got := utils.CalculatePercentageDifference(newNum, originNum)

	want := utils.Round4Decimal(-13.652275379229861)

	if got != want {
		t.Errorf(`got "%f", want "%f"`, got, want)
	}
}

func TestRound4Decimal(t *testing.T) {

	num := -13.652275379229861

	got := utils.Round4Decimal(num)

	want := -13.6523

	if got != want {
		t.Errorf(`got "%f", want "%f"`, got, want)
	}
}
