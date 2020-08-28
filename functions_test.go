package main

import "testing"

func TestCompInterest(t *testing.T) {
	// TODO: add testing for CompoundInterest()
	var (
		p float64 = 1000
		apr float64 = 0.33
		n float64 = 3
		y float64 = 8
	)

	result := CompoundInterest(p, apr, n, y)

	if result != 12239.16 {
		t.Errorf("CompoundInterest(1000, 0.33, 3, 8) = %.2f; wanted : 12239.16", result)
	}
}

func TestCompInterestYearly(t *testing.T) {
	// TODO: add testing for CompoundInterestYearly()
	var (
		p float64 = 1000
		apr float64 = 0.33
		y float64 = 8
	)

	result := CompoundInterestYearly(p, apr, y)

	if result != 9790.69 {
		t.Errorf("CompoundInterestYearly(1000, 0.33, 8) = %.2f; wanted : 9790.69", result)
	}
}

func TostContCompIterest(t *testing.T) {
	// TODO: add testing for ContinuousCompoundInterest()
	var (
		p float64 = 1000
		apr float64 = 0.33
		y float64 = 8
	)

	result := ContinuousCompoundInterest(p, apr, y)

	if result != 14013.20 {
		t.Errorf("ContinuousCompoundTest(1000, 0.33, 8) = %.2f; wanted : 14013.20", result)
	}
}

func TestSavePlan(t *testing.T) {
	var (
		pmt float64 = 40
		apr float64 = 0.33
		n float64 = 4
		y float64 = 20
	)

	result := SavingsPlanReg(pmt, apr, n, y)

	// if result != 4881.55403 {
	if result != 274835.59 {
		t.Errorf("SavingsPlanReg(1000, 0.33, 4, 20) = %.2f; wanted : 274835.59", result)
	}
}