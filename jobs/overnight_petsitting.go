package jobs

import (
	"fmt"
	"time"

	"github.com/cloudtriquetra/payout/db"
	"github.com/cloudtriquetra/payout/employee"
	"github.com/cloudtriquetra/payout/utils"
)

type EffortEntryOvernightPetSitting struct {
	Effort
	PetName string
}

func newEffortInputForOvernightPetSitting(empName string, effortDate string, description string, petName string) (*EffortEntryOvernightPetSitting, error) {

	return &EffortEntryOvernightPetSitting{
		Effort: Effort{
			EmployeeName: empName,
			Description:  description,
			EffortDate:   effortDate,
			EffortType:   "Overnight Pet Sitting",
			Amount:       jobRates["overnight_petsitting"]},
		PetName: petName,
	}, nil

}

func PostEffortInputOvernightPetSitting() EffortEntryOvernightPetSitting {
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

	effortEntry, err := newEffortInputForOvernightPetSitting(empName, effortDate, description, petName)
	if err != nil {
		fmt.Println("Error with Effort Entry")
		panic(err)
	}

	effortEntry.save()
	return *effortEntry
}

func (e EffortEntryOvernightPetSitting) save() {
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
	fmt.Println("Effort Entry for Overnight Pet Sitting has been saved successfully")

}
