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

func GetUserMetrics() (weight, height float64, err error) {
	weight, err = getUserInput(weightPrompt)
	if err != nil {
		return
	}
	height, err = getUserInput(heightPrompt)
	if err != nil {
		return
	}
	return
}

////////////////////////
//     user input     //
////////////////////////

func getUserInput(promptText string) (value float64, err error) {
	fmt.Print(promptText)
	// on enter submit the input
	userInput, _ := reader.ReadString('\n')
	// remove the extra '\n' added by enter
	userInput = strings.Replace(userInput, "\n", "", -1)
	// convert string to float
	value, err = strconv.ParseFloat(userInput, 64)
	return
}
