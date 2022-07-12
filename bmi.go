package main

////////////////////////
//  BMI calculations  //
////////////////////////

func calculateBMI(weight, height float64) float64 {
	// convert height in meters
	heightInMeter := height / 100
	// bmi formula = mass / height**2 [kg/m**2]
	bmi := weight / (heightInMeter * heightInMeter)
	return bmi
}

// // ANOTHER WAY USING CHANNELS
// func calculateBMIWithChannels(weight, height float64, ch chan float64) {
// 	// convert height in meters
// 	heightInMeter := height / 100
// 	// bmi formula = mass / height**2 [kg/m**2]
// 	bmi := weight / (heightInMeter * heightInMeter)
// 	ch <- bmi
// }