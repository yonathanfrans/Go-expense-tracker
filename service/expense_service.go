package service

import (
	"encoding/csv"
	"errors"
	"example/expense-tracker/model"
	"example/expense-tracker/storage"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	ErrExpenseNotFound = errors.New("expense not found")
	ErrEmptyDescription = errors.New("description cannot be empty")
	ErrInvalidAmount = errors.New("amount must be greater than 0")
	ErrInvalidMonth = errors.New("month must be between 1 and 12")	
	ErrEmptyCategory = errors.New("category cannot be empty")
)

func findExpenseIndexByID (expenses []model.Expense, id int) int {
	for i := range expenses {
		if expenses[i].ID == id {
			return i
		}
	}

	return -1
}

func AddExpense(description string, amount float64, category string) (int, error) {
	if strings.TrimSpace(description) == "" {
		return 0, ErrEmptyDescription
	}
	
	if amount <= 0 {
		return 0, ErrInvalidAmount
	}

	if strings.TrimSpace(category) == "" {
		return 0, ErrEmptyCategory
	}

	expenses, err := storage.LoadExpenses()
	if err != nil {
		return 0, err
	}

	maxID := 0
	for _, expense := range expenses {
		if expense.ID > maxID {
			maxID = expense.ID
		}
	}
	nextID := maxID + 1

	newExpense := model.Expense{
		ID: nextID,
		Description: description,
		Amount: amount,
		Category: category,
		Date: time.Now(),
	}

	expenses = append(expenses, newExpense)

	err = storage.SaveExpenses(expenses)
	if err != nil {
		return 0, err
	}

	checkBudgetWarning(int(time.Now().Month()))

	return nextID, nil
}

func ListExpenses(category string) error {
	expenses, err := storage.LoadExpenses()
	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found")
		return nil
	}

	hasDisplayed := false
	for _, expense := range expenses {
		if category != "" && !strings.EqualFold(expense.Category, category) {
			continue
		}

		fmt.Println("==========")
		fmt.Printf("ID: %d\nDescription: %s\nAmount: %.2f\nCategory: %s\nDate: %s\n", 
				expense.ID, 
				expense.Description,
				expense.Amount,
				expense.Category,
				expense.Date.Format("2006-01-02"),
		)
		hasDisplayed = true
	}

	if !hasDisplayed {
		fmt.Printf("No expenses found for category '%s'\n", category)
	}

	return nil
}

func UpdateExpense(id int, description string, amount float64, category string) error {
	if strings.TrimSpace(description) == "" {
		return ErrEmptyDescription
	}

	if amount <= 0 {
		return ErrInvalidAmount
	}

	if strings.TrimSpace(category) == "" {
		return ErrEmptyCategory
	}

	expenses, err := storage.LoadExpenses()
	if err != nil {
		return err
	}
	
	idx := findExpenseIndexByID(expenses, id)
	if idx == -1 {
		return ErrExpenseNotFound
	}

	expenses[idx].Description = description
	expenses[idx].Amount = amount
	expenses[idx].Category = category

	return storage.SaveExpenses(expenses)
}

func DeleteExpense(id int) error {
	expenses, err := storage.LoadExpenses()
	if err != nil {
		return err
	}

	idx := findExpenseIndexByID(expenses, id)
	if idx == -1 {
		return ErrExpenseNotFound
	}

	expenses = append(expenses[:idx], expenses[idx+1:]...)

	return storage.SaveExpenses(expenses)
}

func SummaryExpenses(month int) (float64, error){
	expenses, err := storage.LoadExpenses()
	if err != nil {
		return 0, err
	}

	if month < 0 || month > 12 {
		return 0, ErrInvalidMonth
	}

	var total float64

	for _, expense := range expenses {
		if month != 0 && int(expense.Date.Month()) != month {
			continue
		}
		total += expense.Amount
	}

	return total, nil
}

func SetMonthlyBudget(month int, amount float64) error {
	if month < 1 || month > 12 {
		return ErrInvalidMonth
	}

	if amount <= 0 {
		return ErrInvalidAmount
	}

	budgets, err := storage.LoadBudgets()
	if err != nil {
		return err
	}

	found := false 
	for i := range budgets {
		if budgets[i].Month == month {
			budgets[i].Amount = amount
			found = true
			break
		}
	}

	if !found {
		budgets = append(budgets, model.Budget{Month: month, Amount: amount})
	}

	return storage.SaveBudgets(budgets)
}

func checkBudgetWarning(month int) {
	budgets, err := storage.LoadBudgets()
	if err != nil || len(budgets) == 0 {
		return
	}

	var targetBudget float64
	for _, b := range budgets {
		if b.Month == month {
			targetBudget = b.Amount
			break
		}
	}

	if targetBudget == 0 {
		return
	}

	totalSpent, _ := SummaryExpenses(month)

	if totalSpent > targetBudget {
		fmt.Printf("\n Warning:Pengeluaran Anda bulan ini (Rp.%.2f) telah MELEBIHI batas budget (Rp.%.2f)!\n", totalSpent, targetBudget)
	}
}

func ExportToCSV() error {
	expenses, err := storage.LoadExpenses()
	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		return errors.New("tidak ada data expense untuk diexport")
	}

	file, err := os.Create("expenses.csv")
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "Description", "Amount", "Category", "Date"}
	err = writer.Write(header)
	if err != nil {
		return err
	}

	for _, expense := range expenses {
		row := []string{
			strconv.Itoa(expense.ID),
			expense.Description,
			fmt.Sprintf("%.2f", expense.Amount),
			expense.Category,
			expense.Date.Format("2006-01-02"),
		}

		err = writer.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}