package main

import (
	"fmt"
	"math"
)

func main() {

	/*
		var investAmount = 1000
		var expectedReturnRate = 5.5
		var years = 10
		var futureValue = float64(investAmount) * math.Pow((1+expectedReturnRate/100), float64(years))
		fmt.Println(futureValue)
	*/
	//------------------------------------------------------
	// we could uses explicit type assignment insted of doing the type conversion later
	// var investAmount float64 = 1000
	// internaly 1000 will be stored as 1000.0
	// and using that we can get rid of the type conversions
	//------------------------------------------------------
	/*
		var investAmount float64 = 1000
		var expectedReturnRate = 5.5
		var years float64 = 10
		var futureValue = investAmount * math.Pow((1+expectedReturnRate/100), years)
		fmt.Println(futureValue)
	*/
	//------------------------------------------------------
	// Alternative variable declaration
	// Here we use := to declare and assign a value to the variable and we can't use explicit type here
	// we can declare multiple varables on the same line and they don't need to be of the same type
	var investAmount, years float64 = 1000, 10
	expectedReturnRate := 5.5
	futureValue := investAmount * math.Pow((1+expectedReturnRate/100), years)
	fmt.Println(futureValue)

}
