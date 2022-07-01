package main

import "github.com/akshayGudhate/bmi/info"

func main() {
	// print welcome message
	info.PrintWelcomeMessage()

	// get user metrics data
	weight, height := info.GetUserMetrics()

	// bmi calculations
	bmi := calculateBMI(weight, height)

	// print bmi message
	info.PrintBMIMessage(bmi)
}
