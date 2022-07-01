package info

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

////////////////////////
//    user metrics    //
////////////////////////

func GetUserMetrics() (weight float64, height float64) {
	// get user weight
	weight = getUserInput(weightPrompt)
	// get user height
	height = getUserInput(weightPrompt)
	return
}

////////////////////////
//     user input     //
////////////////////////

func getUserInput(promptText string) (enteredValue float64) {
	fmt.Print(promptText)
	// on enter submit the input
	userInput, _ := reader.ReadString('\n')
	// remove the extra '\n' added by enter
	userInput = strings.Replace(userInput, "\n", "", -1)
	// convert string to float
	enteredValue, _ = strconv.ParseFloat(userInput, 64)
	return
}
