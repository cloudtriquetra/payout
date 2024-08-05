package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetSingleUserInput(infoText string) string {
	var userInputValue string
	fmt.Print(infoText)
	fmt.Scanln(&userInputValue)

	return userInputValue

}

func GetMultiUserInput(infoText string) string {
	fmt.Print(infoText)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
