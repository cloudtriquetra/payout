package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DBeffort *sql.DB
var DBexpense *sql.DB
var DBhomecare *sql.DB

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

type Expense struct {
	ExpenseID          int
	EmployeeName       string
	ExpenseDate        string
	ExpenseDescription string
	ExpenseAmount      float64
}

type HomeCare struct {
	HomeCareID          int
	EmployeeName        string
	HomeCareStartTime   string
	HomeCareEndTime     string
	HomeCareDuration    float64
	HomeCareDescription string
	HomeCareAmount      float64
	HomeCarePetName     string
	HomeCarePetType     string
}

func InitDB() {
	var err error
	const errMsg = "Error opening database: "
	DBeffort, err = sql.Open("sqlite3", "effort.db")
	if err != nil {
		panic(errMsg + err.Error())
	}

	DBexpense, err = sql.Open("sqlite3", "expense.db")
	if err != nil {
		panic(errMsg + err.Error())
	}

	DBhomecare, err = sql.Open("sqlite3", "homecare.db")
	if err != nil {
		panic(errMsg + err.Error())
	}

	DBeffort.SetMaxOpenConns(10)
	DBeffort.SetMaxIdleConns(5)
	DBexpense.SetMaxOpenConns(10)
	DBexpense.SetMaxIdleConns(5)
	DBhomecare.SetMaxOpenConns(10)
	DBhomecare.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	effortsqlStmt := `
	CREATE TABLE IF NOT EXISTS efforts (
    effort_id               INTEGER PRIMARY KEY AUTOINCREMENT,
    employee_name           varchar(40) NOT NULL,
    effort_type             varchar(40) NOT NULL,
    effort_date             date NOT NULL,
    start_time              varchar(40) default NA,
    end_time                varchar(40) default NA,
    effort_description      text default NA,
    pet_name                varchar(40)default NA,
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

	homecaresqlStmt := `CREATE TABLE IF NOT EXISTS homecare (
	homecare_id             INTEGER PRIMARY KEY AUTOINCREMENT,
	employee_name           text NOT NULL,
	homecare_start_time     date NOT NULL,
	homecare_end_time	    date NOT NULL,
	homecare_duration       numeric(10,2) NOT NULL,
	homecare_description    text,
	homecare_amount         numeric(10,2),
	homecare_pet_name       text NOT NULL,
	homecare_pet_type       text NOT NULL check(homecare_pet_type = "dog" or homecare_pet_type = "cat")
	);`

	_, err := DBeffort.Exec(effortsqlStmt)
	if err != nil {
		panic("Effort DB Creation Failed" + err.Error())
	}

	_, err = DBexpense.Exec(expensesqlStmt)
	if err != nil {
		panic("Expense DB Creation Failed" + err.Error())
	}

	_, err = DBhomecare.Exec(homecaresqlStmt)
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
