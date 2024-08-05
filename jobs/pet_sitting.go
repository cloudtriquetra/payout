package jobs

import (
	"fmt"
	"os"
	"time"

	"github.com/cloudtriquetra/payout/db"
	"github.com/cloudtriquetra/payout/employee"
	"github.com/cloudtriquetra/payout/utils"
)

type EffortEntryPetSitting struct {
	EmployeeName   string
	EffortDate     string
	StartTimeStamp time.Time
	EndTimeStamp   time.Time
	DurationInHour float64
	Description    string
	EffortType     string
	EffortRate     float64
	Amount         float64
	PetName        string
}

func NewEffortInputForPetSitting(empName string, startTimeStamp time.Time, endTimeStamp time.Time, duration float64, description string, effortDate string, petName string) (*EffortEntryPetSitting, error) {

	return &EffortEntryPetSitting{
		EmployeeName:   empName,
		StartTimeStamp: startTimeStamp,
		EndTimeStamp:   endTimeStamp,
		DurationInHour: duration,
		Description:    description,
		EffortDate:     effortDate,
		EffortType:     "Pet Sitting",
		EffortRate:     jobRates["pet_sitting"],
		Amount:         jobRates["pet_sitting"] * duration,
		PetName:        petName,
	}, nil
}

func PostEffortInputPetSitting() EffortEntryPetSitting {
	empName, err := employee.GetEmployeeName()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	var startDate, startTime string = utils.GetSingleUserInput("Enter Date for Pet Sitting (DD-MM-YYYY):"),
		utils.GetSingleUserInput("Enter Start Time for Pet Sitting (HH:MM):")

	startTimeStamp, err := utils.GetTimeStamp(startDate, startTime)
	if err != nil {
		fmt.Println("Error with Start Date/Time: ")
		panic(err)
	}

	var endDate, endTime string = startDate,
		utils.GetSingleUserInput("Enter End Time for Pet Sitting (HH:MM): ")

	endTimeStamp, err := utils.GetTimeStamp(endDate, endTime)
	if err != nil {
		fmt.Println("Error with End Date/Time: ")
		panic(err)
	}

	var petName string = utils.GetSingleUserInput("Enter Pet Name:")

	if petName == "" {
		panic("Pet Name is required")
	}

	var duration float64 = utils.GetDuration(
		startTimeStamp,
		endTimeStamp,
	)

	effortDate := startDate

	var description string = utils.GetMultiUserInput("Enter Additional Note / Description for Pet Sitting (OPTIONAL): ")

	effortEntry, err := NewEffortInputForPetSitting(empName, startTimeStamp, endTimeStamp, duration, description, effortDate, petName)
	if err != nil {
		fmt.Println("Error with Effort Entry")
		panic(err)
	}

	effortEntry.Save()
	return *effortEntry

}

func (e EffortEntryPetSitting) Save() {
	// Save EffortEntryPetSitting to DB
	query := `
	INSERT INTO efforts (employee_name, effort_type, effort_date, start_time, end_time, effort_description, duration, cost, pet_name) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`

	stmt, err := db.DBeffort.Prepare(query)

	if err != nil {
		fmt.Println("Error with DB Prepare Statement")
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.EmployeeName, e.EffortType, e.EffortDate, e.StartTimeStamp, e.EndTimeStamp, e.Description, e.DurationInHour, e.Amount, e.PetName)
	if err != nil {
		fmt.Println("Error with DB Exec Statement")
		panic(err)
	}
	result.LastInsertId()

}
