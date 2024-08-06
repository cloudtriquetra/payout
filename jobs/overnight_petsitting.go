package jobs

import (
	"fmt"
	"os"
	"time"

	"github.com/cloudtriquetra/payout/db"
	"github.com/cloudtriquetra/payout/employee"
	"github.com/cloudtriquetra/payout/utils"
)

type EffortEntryOvernightPetSitting struct {
	EmployeeName string
	EffortDate   string
	Description  string
	EffortType   string
	Amount       float64
	PetName      string
}

func NewEffortInputForOvernightPetSitting(empName string, effortDate string, description string, petName string) (*EffortEntryOvernightPetSitting, error) {

	return &EffortEntryOvernightPetSitting{
		EmployeeName: empName,
		Description:  description,
		EffortDate:   effortDate,
		EffortType:   "Overnight Pet Sitting",
		PetName:      petName,
		Amount:       jobRates["overnight_petsitting"],
	}, nil

}

func PostEffortInputOvernightPetSitting() EffortEntryOvernightPetSitting {
	empName, err := employee.GetEmployeeName()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var effortDate string = utils.GetSingleUserInput("Enter Start Date for Overnight Petsitting (DD-MM-YYYY):")

	_, err = time.Parse("01-02-2006", effortDate)
	if err != nil {
		panic("Kindly enter a valid date")
	}

	var petName string = utils.GetSingleUserInput("Enter Pet Name:")

	if petName == "" {
		panic("Pet Name is required")
	}

	var description string = utils.GetMultiUserInput("Enter Additional Note / Description for Overnight Pet Sitting (OPTIONAL): ")
	if description == "" {
		description = "NA"
	}

	effortEntry, err := NewEffortInputForOvernightPetSitting(empName, effortDate, description, petName)
	if err != nil {
		fmt.Println("Error with Effort Entry")
		panic(err)
	}

	effortEntry.Save()
	return *effortEntry
}

func (e EffortEntryOvernightPetSitting) Save() {
	query := `
	INSERT INTO efforts (employee_name, effort_type, effort_date, effort_description, cost, pet_name) 
	VALUES (?, ?, ?, ?, ?, ?);`

	stmt, err := db.DBeffort.Prepare(query)

	if err != nil {
		fmt.Println("Error with DB Prepare Statement")
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.EmployeeName, e.EffortType, e.EffortDate, e.Description, e.Amount, e.PetName)
	if err != nil {
		fmt.Println("Error with DB Exec Statement")
		panic(err)
	}
	result.LastInsertId()

}
