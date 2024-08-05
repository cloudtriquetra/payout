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
	);

CREATE TABLE IF NOT EXISTS expenses (
	expense_id              INTEGER PRIMARY KEY AUTOINCREMENT,
	employee_name           text NOT NULL,
	expense_date            date NOT NULL,
	expense_description     text NOT NULL,
	expense_amount          numeric(10,2) NOT NULL
	);