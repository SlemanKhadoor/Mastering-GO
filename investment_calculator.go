package main

import (
	"fmt"
	"math"
)

func main() {
	// we could uses explicit type assignment insted of doing the type conversion later
	// var investAmount float64 = 1000
	// internaly 1000 will be stored as 1000.0
	// and using that we can get rid of the type conversions
	//------------------------------------------------------
	/*
		var investAmount = 1000
		var expectedReturnRate = 5.5
		var years = 10
		var futureValue = float64(investAmount) * math.Pow((1+expectedReturnRate/100), float64(years))
		fmt.Println(futureValue)
	*/
	var investAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10
	var futureValue = investAmount * math.Pow((1+expectedReturnRate/100), years)
	fmt.Println(futureValue)
}
