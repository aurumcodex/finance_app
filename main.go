package main

import (
	"fmt"
	"math"
	"strconv"
)

// func

func main() {
	var (
		pmtS string
		pS   string
		aprS string
		nS   string
		tS   string
	)

	num1 := CompoundInterest(1000, 0.42, 2, 10)
	num2 := CompoundInterestYearly(1000, 0.42, 10)
	num3 := ContinuousCompoundInterest(1000, 0.42, 10)
	num4 := SavingsPlanReg(20, 0.42, 1, 10)

	fmt.Println("grug grug grug")
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)

	fmt.Println("enter payment number")
	fmt.Scanln(&pmtS)
	pmt, _ := strconv.ParseFloat(pmtS, 64)
	fmt.Println("you entered: ", pmt)

	fmt.Println("enter principal number")
	fmt.Scanln(&pS)
	p, _ := strconv.ParseFloat(pS, 64)
	fmt.Println("you entered: ", p)

	fmt.Println("enter APR number")
	fmt.Scanln(&aprS)
	apr, _ := strconv.ParseFloat(aprS, 64)
	fmt.Println("you entered: ", apr)

	fmt.Println("enter period frequency number")
	fmt.Scanln(&nS)
	n, _ := strconv.ParseFloat(nS, 64)
	fmt.Println("you entered: ", n)

	fmt.Println("enter number of years")
	fmt.Scanln(&tS)
	t, _ := strconv.ParseFloat(tS, 64)
	fmt.Println("you entered: ", t)

	fmt.Println("your compounded interest is: ", CompoundInterest(p, apr, n, t))
	fmt.Println("your yearly compounded interest is: ", CompoundInterestYearly(p, apr, t))
	fmt.Println("your continuous compounded interest is: ", ContinuousCompoundInterest(p, apr, t))
	fmt.Println("your savings plan would be: ", SavingsPlanReg(pmt, apr, n, t))

	num := 274835.585090
	fmt.Println("rounded decimal: ", math.Round(num*100)/100)
}
