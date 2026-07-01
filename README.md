# Expense Tracker CLI

A simple Command Line Interface (CLI) application built with Go to track daily expense. This project is based on the Expense Tracker project from roadmap.sh and was developed as part of my Go backend learning journey.

---

## Features
- Add expense
- Update expense
- Delete expense
- List all expenses
- Filter expenses by category
- Show total expense summary
- Show monthly expense summary
- Set monthly budget
- Budget warning when expenses exceed the budget
- Export expenses to CSV

---

## Tech Stack

* Go
* Standard Library
* JSON Storage
* CSV Export

---

## Project Structure
```
expense-tracker/
│
├── model/
│   └── expense.go
│
├── service/
│   └── expense_service.go
│
├── storage/
│   └── json_storage.go
│
├── main.go
├── go.mod
└── README.md
```

---

## Getting Started

Clone the repository:

```bash
git clone https://github.com/yourusername/expense-tracker-go.git
```

Move into the project

```bash
cd expense-tracker-go
```

Run

```bash
go run .
```

---

## Usage

### Add Expense

```bash
go run . add --description "Lunch" --amount 25000 --category Food
```

### Update Expense

```bash
go run . update --id 1 --description "Dinner" --amount 50000 --category Food
```

### Delete Expense

```bash
go run . delete --id 1
```

### List Expenses

```bash
go run . list
```

### Filter by Category

```bash
go run . list --category Food
```

### Summary

```bash
go run . summary
```

### Monthly Summary

```bash
go run . summary --month 8
```

### Set Monthly Budget

```bash
go run . set-budget --month 8 --amount 1000000
```

### Export CSV

```bash
go run . export
```

---

## Concepts Practiced

- Struct
- Functions
- Packages
- Error Handling
- JSON Encoding/Decoding
- File Handling
- CLI Argument Parsing
- CSV Export
- Clean Project Structure

---

## Future Improvements

- Support multiple currencies
- Expense statistics
- Better CLI help command
- Unit tests
- SQLite storage
- Colored terminal output

---

## Acknowledgements

This project is based on the Expense Tracker challenge from roadmap.sh.

## Project Source

https://roadmap.sh/projects/expense-tracker
