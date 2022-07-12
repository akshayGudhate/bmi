package main

import "github.com/akshayGudhate/bmi/info"

func main() {
	// print welcome message
	info.PrintWelcomeMessage()

	// get user metrics data
	weight, height, err := info.GetUserMetrics()
	// if error
	if err != nil {
		info.PrintInvalidInputMessage()
		return
	}

	// bmi calculations
	bmi := calculateBMI(weight, height)

	// // ANOTHER WAY USING CHANNELS
	// ch := make(chan float64)
	// go calculateBMIWithChannels(weight, height, ch)
	// bmiWithChannel := <-ch
	// info.PrintBMIMessage(bmiWithChannel)

	// print end message
	info.PrintEndingMessage()

	// print bmi message
	info.PrintBMIMessage(bmi)
}
