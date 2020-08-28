package main

import "math"

func CompoundInterest(p, apr, n, t float64) float64 {
	parens := 1.0 + (apr / n)
	exp := n * t

	return math.Round((p * math.Pow(parens, exp)) * 100) / 100

	// return math.Round((p * math.Pow((1.0 + (apr / n)), t)) * 100) / 100
}

func CompoundInterestYearly(p, apr, t float64) float64 {
	return math.Round((p * math.Pow((1.0 + apr), t)) * 100) / 100
}

func ContinuousCompoundInterest(p, apr, t float64) float64 {
	return math.Round((p * math.Pow(math.E, (apr*t))) * 100) / 100
}

func SavingsPlanReg(pmt, apr, n, t float64) float64 {
	top := math.Pow((1.0+(apr/n)), n*t) - 1.0
	bottom := apr / n

	return math.Round((pmt * top / bottom) * 100) / 100
}
