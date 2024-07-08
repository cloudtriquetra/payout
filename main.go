package main

import (
	"fmt"
	"time"

	"github.com/cloudtriquetra/payout/utils"
)

func main() {
	var empName string = utils.GetUserInput("Enter Employee Name:")
	var startDate, startTime string = utils.GetUserInput("Enter Start Date (DD-MM-YYYY):"),
		utils.GetUserInput("Enter Start Time (HH:MM):")

	startTimeStamp, err := utils.GetTimeStamp(startDate, startTime)
	if err != nil {
		fmt.Println("Error with Start Date/Time: ")
		panic(err)
	}

	var endDate, endTime string = utils.GetUserInput("Enter End Date (DD-MM-YYYY):"),
		utils.GetUserInput("Enter End Time (HH:MM): ")

	endTimeStamp, err := utils.GetTimeStamp(endDate, endTime)
	if err != nil {
		fmt.Println("Error with End Date/Time: ")
		panic(err)
	}

	var duration time.Duration = utils.GetDuration(
		startTimeStamp,
		endTimeStamp,
	)

	effortEntry, err := utils.NewEffortEntry(empName, startTimeStamp, endTimeStamp, duration)
	if err != nil {
		fmt.Println("Error with Effort Entry")
		panic(err)
	}

	effortEntry.RecordEffortToFile("a.csv")
}
