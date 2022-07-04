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

func GetUserMetrics() (weight, height float64) {
	weight = getUserInput(weightPrompt)
	height = getUserInput(heightPrompt)
	return
}

////////////////////////
//     user input     //
////////////////////////

func getUserInput(promptText string) (value float64) {
	fmt.Print(promptText)
	// on enter submit the input
	userInput, _ := reader.ReadString('\n')
	// remove the extra '\n' added by enter
	userInput = strings.Replace(userInput, "\n", "", -1)
	// convert string to float
	value, _ = strconv.ParseFloat(userInput, 64)
	return
}
