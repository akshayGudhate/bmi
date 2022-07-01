package info

import "fmt"

////////////////////////
//       welcome      //
////////////////////////

func PrintWelcomeMessage() {
	fmt.Println(separator)
	fmt.Println(mainTitle)
	fmt.Println(separator)
}

////////////////////////
//      bmi case      //
////////////////////////

func PrintBMIMessage(bmi float64) {
	// show user his bmi as per type
	if underWeight <= bmi && bmi < normalWeight {
		fmt.Printf(printUnderWeightMessage, bmi)
	} else if normalWeight <= bmi && bmi < overWeight {
		fmt.Printf(printNormalWeightMessage, bmi)
	} else if overWeight <= bmi && bmi < obeseWeight {
		fmt.Printf(printOverWeightMessage, bmi)
	} else if bmi >= obeseWeight {
		fmt.Printf(printObeseWeightMessage, bmi)
	} else {
		fmt.Printf(printOtherWeightMessage, bmi)
	}
}
