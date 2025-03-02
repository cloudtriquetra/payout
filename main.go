package main

import (
	"fmt"
	"os"

	"github.com/cloudtriquetra/payout/db"
	"github.com/cloudtriquetra/payout/jobs"
	"github.com/cloudtriquetra/payout/utils"
)

func main() {
	db.InitDB()
	for {
		choice := utils.GetSingleUserInput(`Enter Job/Expense Type:
	1. Hotel Shift
	2. Pet Sitting
	3. Cat Visit
	4. Overnight Hotel Shift
	5. Overnight Pet Sitting
	6. Cat at Sitter's Home
	7. Dog at Sitter's Home
	8. Uber / Expense
	9. Exit
Your Choice: `)

		switch choice {
		case "1":
			jobs.PostEffortInputHotel()

		case "2":
			jobs.PostEffortInputPetSitting()

		case "3":
			jobs.PostEffortInputCatVisit()

		case "4":
			jobs.PostEffortInputOvernightHotel()

		case "5":
			jobs.PostEffortInputOvernightPetSitting()

		case "8":
			jobs.PostExpense()

		case "9":
			os.Exit(0)

		default:
			fmt.Println("Invalid Choice")
		}
	}
}
