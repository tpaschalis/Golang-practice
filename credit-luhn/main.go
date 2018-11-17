package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

    // Initial read of card number
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Card Number : ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSuffix(input, "\n")

	if _, err := strconv.ParseInt(input, 10, 64); err != nil {
		fmt.Printf("Input is not numerical\n")
		os.Exit(1)
	}

    // Detection of card type (AMEX/Visa/Mastercard)
	firstTwo, err := strconv.Atoi(input[0:2])
    if err != nil {
        panic(err)
    }
	currentCard := cardType(firstTwo, input)

	// Calculate Luhn's checksum
	sum := 0
	doubleFlag := false
	for i := len(input); i >= 1; i = i - 1 {
		currentPlace := input[i-1 : i]
		currentDigit, _ := strconv.Atoi(currentPlace)
		if doubleFlag {
			sum = sum + addDigits(strconv.Itoa(currentDigit*2))
		} else {
			sum = sum + currentDigit
		}
		doubleFlag = swapBoolean(doubleFlag)
	}

	// Checksum validation
	if (sum)%10 == 0 {
		fmt.Println("This", currentCard, "card is valid")
	} else {
		fmt.Println("This", currentCard, "card is invalid, sorry!")
	}

}

func cardType(firstTwo int, input string) string {
	cardType := ""
	switch {
	case (firstTwo == 34) && len(input) == 15:
		cardType = "American Express"
	case (firstTwo == 37) && len(input) == 15:
		cardType = "American Express"
	case (firstTwo <= 55) && (firstTwo >= 51) && len(input) == 16:
		cardType = "MasterCard"
	case (firstTwo <= 49) && (firstTwo >= 40) && len(input) == 16:
		cardType = "Visa"
	default:
		if cardType == "" {
			fmt.Printf("Invalid card type\n")
			os.Exit(1)
		}
	}

	return cardType
}

func addDigits(num string) int {

	length := len(num)
	sum := 0
	for i := 0; i < length; i++ {
		temp, err := strconv.Atoi(num[i : i+1])
		if err != nil {
			panic(err)
		}
		sum = sum + temp
	}
	if sum > 9 {
		return (addDigits(strconv.Itoa(sum)))
	}

	return sum
}

func swapBoolean(v bool) bool {
	if v {
		return false
	} else {
		return true
	}
}

// Possible improvements list ^^
//
// Limit usage of strconv
// Add recognition of more card types
// Decide if string handling is good,
// or it's better to use some long int
//
