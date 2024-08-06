package main

import (
	"github.com/cloudtriquetra/payout/db"
	"github.com/cloudtriquetra/payout/jobs"
	"github.com/cloudtriquetra/payout/utils"
)

func main() {
	db.InitDB()
	choice := utils.GetSingleUserInput(`Enter Job/Expense Type:
	1. Hotel Shift
	2. Pet Sitting
	3. Cat Visit
	4. Overnight Hotel Shift
	5. Overnight Pet Sitting
	6. Cat at Sitter Home
	7. Dog at Sitter Home
	8. Uber / Expense
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
	}
}
