package main

import (
	"fmt"

	"github.com/cloudtriquetra/payout/utils"
)

func main() {
	var empName string = utils.GetUserInput("Enter Employee Name:")
	var startDate, startTime, endDate, endTime string = utils.GetUserInput("Enter Start Date (DD-MM-YYYY):"),
		utils.GetUserInput("Enter Start Time (HH:MM):"),
		utils.GetUserInput("Enter End Date (DD-MM-YYYY):"),
		utils.GetUserInput("Enter End Time (HH:MM):")

	startTimeStamp := utils.GetTimeStamp(startDate, startTime)
	endTimeStamp := utils.GetTimeStamp(endDate, endTime)

	var duration string = utils.GetDuration(
		startTimeStamp,
		endTimeStamp,
	).String()

	effortEntry := fmt.Sprintln(empName, startTimeStamp.String(), endTimeStamp.String(), duration)
	fmt.Println(effortEntry)
}
