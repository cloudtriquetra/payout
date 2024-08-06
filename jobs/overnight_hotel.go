package jobs

import (
	"fmt"
	"time"

	"github.com/cloudtriquetra/payout/db"
	"github.com/cloudtriquetra/payout/employee"
	"github.com/cloudtriquetra/payout/utils"
)

type EffortEntryOvernightHotel struct {
	Effort
}

func newEffortInputForOvernightHotel(empName string, effortDate string, description string) (*EffortEntryOvernightHotel, error) {

	return &EffortEntryOvernightHotel{
		Effort: Effort{
			EmployeeName: empName,
			Description:  description,
			EffortDate:   effortDate,
			EffortType:   "Overnight Hotel",
			Amount:       jobRates["overnight_hotel"]},
	}, nil

}

func PostEffortInputOvernightHotel() EffortEntryOvernightHotel {
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

	var effortDate string = utils.GetSingleUserInput("Enter Start Date for Overnight Hotel Shift (DD-MM-YYYY):")

	_, err = time.Parse("01-02-2006", effortDate)
	if err != nil {
		panic("Kindly enter a valid date")
	}

	var description string = utils.GetMultiUserInput("Enter Additional Note / Description for Overnight Hotel Shift (OPTIONAL): ")
	if description == "" {
		description = "NA"
	}

	effortEntry, err := newEffortInputForOvernightHotel(empName, effortDate, description)
	if err != nil {
		fmt.Println("Error with Effort Entry")
		panic(err)
	}

	effortEntry.save()
	return *effortEntry
}

func (e EffortEntryOvernightHotel) save() {
	query := `
	INSERT INTO efforts (employee_name, effort_type, effort_date, effort_description, cost) 
	VALUES (?, ?, ?, ?, ?);`

	stmt, err := db.DBeffort.Prepare(query)

	if err != nil {
		fmt.Println("Error with DB Prepare Statement")
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.EmployeeName, e.EffortType, e.EffortDate, e.Description, e.Amount)
	if err != nil {
		fmt.Println("Error with DB Exec Statement")
		panic(err)
	}
	result.LastInsertId()
	fmt.Println("Effort Entry for Overnight Hotel Shift has been saved successfully")

}
