package info

/////////////////////////
//      constants      //
/////////////////////////

const (
	// bmi type
	underWeight  = 01.0
	normalWeight = 18.5
	overWeight   = 25.0
	obeseWeight  = 30.0

	// user input
	separator    = "--------------------------------------------------------------"
	mainTitle    = "*********************** BMI Calculator ***********************"
	weightPrompt = "Please enter your weight (kg): "
	heightPrompt = "Please enter your height (cm): "

	// bmi output
	messageUnderWeight  = "\nYour BMI is: %.2f and you are underweight person.\n"
	messageNormalWeight = "\nYour BMI is: %.2f and you have perfectly normal weight.\n"
	messageOverWeight   = "\nYour BMI is: %.2f and you are overweight person.\n"
	messageObeseWeight  = "\nYour BMI is: %.2f and you are obese person.\n"
	messageOtherWeight  = "\nYour BMI is: %.2f.\n"
)
