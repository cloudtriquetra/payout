package jobs

import (
	"fmt"
	"time"

	"github.com/cloudtriquetra/payout/db"
	"github.com/cloudtriquetra/payout/employee"
	"github.com/cloudtriquetra/payout/utils"
)

type EffortEntryPetSitting struct {
	Effort
	StartTimeStamp time.Time
	EndTimeStamp   time.Time
	DurationInHour float64
	PetName        string
}

func newEffortInputForPetSitting(empName string, startTimeStamp time.Time, endTimeStamp time.Time, duration float64, description string, effortDate string, petName string) (*EffortEntryPetSitting, error) {
	// New Effort Entry for Pet Sitting
	return &EffortEntryPetSitting{
		Effort: Effort{
			EmployeeName: empName,
			Description:  description,
			EffortDate:   effortDate,
			EffortType:   "Pet Sitting",
			Amount:       jobRates["pet_sitting"] * duration},
		StartTimeStamp: startTimeStamp,
		EndTimeStamp:   endTimeStamp,
		DurationInHour: duration,
		PetName:        petName,
	}, nil
}

func PostEffortInputPetSitting() EffortEntryPetSitting {
	// Post Effort Entry for Pet Sitting
	var empName string
	var err error

	for {
		empName, err = employee.GetEmployeeName()
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
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
	if description == "" {
		description = "NA"
	}

	effortEntry, err := newEffortInputForPetSitting(empName, startTimeStamp, endTimeStamp, duration, description, effortDate, petName)
	if err != nil {
		fmt.Println("Error with Effort Entry")
		panic(err)
	}

	effortEntry.save()
	return *effortEntry

}

func (e EffortEntryPetSitting) save() {
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
	fmt.Println("Effort Entry for Pet Sitting has been saved successfully")

}
