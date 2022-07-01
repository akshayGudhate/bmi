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

// show user his bmi as per type
func PrintBMIMessage(bmi float64) {

	if underWeight <= bmi && bmi < normalWeight {

		fmt.Printf(messageUnderWeight, bmi)

	} else if normalWeight <= bmi && bmi < overWeight {

		fmt.Printf(messageNormalWeight, bmi)

	} else if overWeight <= bmi && bmi < obeseWeight {

		fmt.Printf(messageOverWeight, bmi)

	} else if bmi >= obeseWeight {

		fmt.Printf(messageObeseWeight, bmi)

	} else {

		fmt.Printf(messageOtherWeight, bmi)
	}
}
