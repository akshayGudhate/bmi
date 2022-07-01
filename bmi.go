package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	fmt.Println(separator)
	fmt.Println(mainTitle)
	fmt.Println(separator)

	////////////////////////
	//       weight       //
	////////////////////////

	fmt.Print(weightPrompt)
	// on enter submit the input
	enteredWeight, _ := reader.ReadString('\n')
	// remove the extra '\n' added by enter
	enteredWeight = strings.Replace(enteredWeight, "\n", "", -1)
	// convert string to float
	weight, _ := strconv.ParseFloat(enteredWeight, 64)

	////////////////////////
	//       height       //
	////////////////////////

	fmt.Print(heightPrompt)
	// on enter submit the input
	enteredHeight, _ := reader.ReadString('\n')
	// remove the extra '\n' added by enter
	enteredHeight = strings.Replace(enteredHeight, "\n", "", -1)
	// convert string to float
	height, _ := strconv.ParseFloat(enteredHeight, 64)

	////////////////////////
	//    calculations    //
	////////////////////////

	// bmi formula = mass/height*height [kg/m**2]
	bmi := (weight / (height * height)) * 10000

	////////////////////////
	//     bmi output     //
	////////////////////////

	// show user his bmi as per type
	if underWeight <= bmi && bmi < normalWeight {
		fmt.Printf(printUnderWeight, bmi)
	} else if normalWeight <= bmi && bmi < overWeight {
		fmt.Printf(printNormalWeight, bmi)
	} else if overWeight <= bmi && bmi < obeseWeight {
		fmt.Printf(printOverWeight, bmi)
	} else if bmi >= obeseWeight {
		fmt.Printf(printObeseWeight, bmi)
	} else {
		fmt.Printf(printOtherWeight, bmi)
	}
}
