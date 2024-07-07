package utils

import "fmt"

func GetUserInput(infoText string) string {
	var userInputValue string
	fmt.Print(infoText)
	fmt.Scan(&userInputValue)

	return userInputValue

}
