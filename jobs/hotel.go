package jobs

import (
	"fmt"
	"os"
	"time"

	"github.com/cloudtriquetra/payout/db"
	"github.com/cloudtriquetra/payout/employee"
	"github.com/cloudtriquetra/payout/utils"
)

type EffortEntryHotel struct {
	EffortID       int
	EmployeeName   string
	EffortDate     string
	StartTimeStamp time.Time
	EndTimeStamp   time.Time
	DurationInHour float64
	Description    string
	EffortType     string
	EffortRate     float64
	Amount         float64
}

func NewEffortInputForHotel(empName string, startTimeStamp time.Time, endTimeStamp time.Time, duration float64, effortDate string, description string) (*EffortEntryHotel, error) {

	return &EffortEntryHotel{
		EmployeeName:   empName,
		StartTimeStamp: startTimeStamp,
		EndTimeStamp:   endTimeStamp,
		DurationInHour: duration,
		Description:    description,
		EffortDate:     effortDate,
		EffortType:     "Hotel / Day Care",
		EffortRate:     jobRates["hotel_shift"],
		Amount:         jobRates["hotel_shift"] * duration,
	}, nil

}

func PostEffortInputHotel() EffortEntryHotel {
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

	var startDate, startTime string = utils.GetSingleUserInput("Enter Date for Hotel/Day Care Shift (DD-MM-YYYY):"),
		utils.GetSingleUserInput("Enter Start Time for Hotel/Day Care Shift (HH:MM):")

	startTimeStamp, err := utils.GetTimeStamp(startDate, startTime)
	if err != nil {
		fmt.Println("Error with Start Date/Time: ")
		os.Exit(0)
	}

	var endDate, endTime string = startDate,
		utils.GetSingleUserInput("Enter End Time for Hotel/Day Care Shift (HH:MM): ")

	endTimeStamp, err := utils.GetTimeStamp(endDate, endTime)
	if err != nil {
		fmt.Println("Error with End Date/Time: ")
		os.Exit(0)
	}

	var duration float64 = utils.GetDuration(
		startTimeStamp,
		endTimeStamp,
	)

	effortDate := startDate

	var description string = utils.GetMultiUserInput("Enter Additional Note / Description for Hotel/Day Care Shift (OPTIONAL): ")
	if description == "" {
		description = "NA"
	}

	effortEntry, err := NewEffortInputForHotel(empName, startTimeStamp, endTimeStamp, duration, effortDate, description)
	if err != nil {
		fmt.Println("Error with Effort Entry")
		panic(err)
	}
	effortEntry.Save()
	return *effortEntry

}

func (e EffortEntryHotel) Save() {
	// Save EffortEntryHotel to DB
	query := `
	INSERT INTO efforts (employee_name, effort_type, effort_date, start_time, end_time, effort_description, duration, cost) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?);`

	stmt, err := db.DBeffort.Prepare(query)

	if err != nil {
		fmt.Println("Error with DB Prepare Statement")
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.EmployeeName, e.EffortType, e.EffortDate, e.StartTimeStamp, e.EndTimeStamp, e.Description, e.DurationInHour, e.Amount)
	if err != nil {
		fmt.Println("Error with DB Exec Statement")
		panic(err)
	}
	result.LastInsertId()
	fmt.Println("Effort Entry for Hotel/Day Care Shift has been saved successfully")

}
