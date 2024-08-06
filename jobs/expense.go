package jobs

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cloudtriquetra/payout/db"
	"github.com/cloudtriquetra/payout/employee"
	"github.com/cloudtriquetra/payout/utils"
)

type Expense struct {
	EmployeeName string
	ExpenseDate  string
	Amount       float64
	Description  string
}

func NewExpense(empName string, expenseDate string, amount float64, description string) (*Expense, error) {

	return &Expense{
		EmployeeName: empName,
		Description:  description,
		ExpenseDate:  expenseDate,
		Amount:       amount,
	}, nil

}

func PostExpense() Expense {
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

	var expenseDate string = utils.GetSingleUserInput("Enter Expense Date (DD-MM-YYYY):")

	_, err = time.Parse("01-02-2006", expenseDate)
	if err != nil {
		panic("Kindly enter a valid date")
	}

	amount, err := strconv.ParseFloat(utils.GetSingleUserInput("Enter Expense Amount: "), 64)
	if err != nil {
		panic("Kindly enter a valid amount" + err.Error())
	}

	if amount == 0.0 {
		panic("Expense Amount is required")
	}

	var description string = utils.GetMultiUserInput("Enter Description for Expense (Mandatory): ")
	if description == "" {
		panic("Expense Description is required")
	}

	expense, err := NewExpense(empName, expenseDate, amount, description)
	if err != nil {
		fmt.Println("Error with Expense Entry")
		panic(err)
	}

	expense.Save()
	return *expense
}

func (e *Expense) Save() {
	query := `INSERT INTO expenses (employee_name, expense_date, expense_amount, expense_description) VALUES (?, ?, ?, ?)`
	stmt, err := db.DBexpense.Prepare(query)
	if err != nil {
		fmt.Println("Error with DB Prepare Statement")
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.EmployeeName, e.ExpenseDate, e.Amount, e.Description)
	if err != nil {
		fmt.Println("Error with DB Exec Statement")
		panic(err)
	}
	result.LastInsertId()
	fmt.Println("Expense Entry Saved Successfully")
}
