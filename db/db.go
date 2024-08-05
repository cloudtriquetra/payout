package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DBeffort *sql.DB
var DBexpense *sql.DB

type Effort struct {
	EffortID       int
	EmployeeName   string
	EffortType     string
	EffortDate     string
	StartTimeStamp sql.NullString
	EndTimeStamp   sql.NullString
	DurationInHour sql.NullFloat64
	Description    sql.NullString
	Amount         float64
	PetName        sql.NullString
}

func InitDB() {
	var err error
	DBeffort, err = sql.Open("sqlite3", "data/effort.db")
	if err != nil {
		panic("Error opening database: " + err.Error())
	}

	DBexpense, err = sql.Open("sqlite3", "data/expense.db")
	if err != nil {
		panic("Error opening database: " + err.Error())
	}

	DBeffort.SetMaxOpenConns(10)
	DBeffort.SetMaxIdleConns(5)
	DBexpense.SetMaxOpenConns(10)
	DBexpense.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	effortsqlStmt := `
	CREATE TABLE IF NOT EXISTS efforts (
    effort_id               INTEGER PRIMARY KEY AUTOINCREMENT,
    employee_name           varchar(40) NOT NULL,
    effort_type             varchar(40) NOT NULL,
    effort_date             date NOT NULL,
    start_time              varchar(40),
    end_time                varchar(40),
    effort_description      text,
    pet_name                varchar(40),
    duration                numeric(10,2),
    cost                    numeric(10,2) NOT NULL
	);`

	expensesqlStmt := `
	CREATE TABLE IF NOT EXISTS expenses (
	expense_id              INTEGER PRIMARY KEY AUTOINCREMENT,
	employee_name           text NOT NULL,
	expense_date            date NOT NULL,
	expense_description     text NOT NULL,
	expense_amount          numeric(10,2) NOT NULL
	);`

	_, err := DBeffort.Exec(effortsqlStmt)
	if err != nil {
		panic("Effort DB Creation Failed" + err.Error())
	}

	_, err = DBexpense.Exec(expensesqlStmt)
	if err != nil {
		panic("Expense DB Creation Failed" + err.Error())
	}

}

func ReadEffortData() ([]Effort, error) {
	query := `SELECT * FROM efforts;`
	rows, err := DBeffort.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var efforts []Effort
	for rows.Next() {
		var effort Effort
		err := rows.Scan(&effort.EffortID, &effort.EmployeeName, &effort.EffortType, &effort.EffortDate, &effort.StartTimeStamp, &effort.EndTimeStamp, &effort.Description, &effort.PetName, &effort.DurationInHour, &effort.Amount)
		if err != nil {
			panic(err)
		}
		efforts = append(efforts, effort)
	}
	return efforts, nil

}
