package info

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
)

////////////////////////
//       welcome      //
////////////////////////

func PrintWelcomeMessage() {
	startTitleMessage := figure.NewColorFigure(startTitle, "", "blue", false)
	startTitleMessage.Print()
	fmt.Println(separator)
}

////////////////////////
//      invalids      //
////////////////////////

func PrintInvalidInputMessage() {
	invalidInputMessage := figure.NewColorFigure(invalidInput, "", "red", false)
	invalidInputMessage.Print()
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
	fmt.Println(separator)
}

////////////////////////
//       ending       //
////////////////////////

func PrintEndingMessage() {
	endTitleMessage := figure.NewColorFigure(endTitle, "", "green", false)
	endTitleMessage.Print()
}
