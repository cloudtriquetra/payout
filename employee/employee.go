package employee

import (
	"errors"
	"slices"
	"strings"

	"github.com/cloudtriquetra/payout/utils"
)

func GetActiveEmployees() []string {
	var employeeNames = []string{
		"prachi",
		"alecia",
		"sagnik",
		"ankita",
		"eray",
	}
	return employeeNames
}

func GetEmployeeName() (string, error) {
	empName := utils.GetSingleUserInput("Enter Employee Name: ")
	if empName == "" {
		return "", errors.New("employee Name is required")
	} else if !slices.Contains(GetActiveEmployees(), strings.ToLower(empName)) {
		return "", errors.New("Invalid employee Name, valid employees are: " + strings.Join(GetActiveEmployees(), ", "))
	} else {
		return empName, nil
	}
}
