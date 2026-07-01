package storage

import (
	"encoding/json"
	"example/expense-tracker/model"
	"os"
)

func LoadExpenses() ([]model.Expense, error) {
	_, err := os.Stat("expenses.json")
	if os.IsNotExist(err) {
		return []model.Expense{}, nil
	}

	file, err := os.ReadFile("expenses.json")
	if err != nil {
		return []model.Expense{}, err
	}

	var expense []model.Expense

	err = json.Unmarshal(file, &expense)
	if err != nil {
		return []model.Expense{}, err
	}

	return expense, nil
}

func SaveExpenses(expense []model.Expense) error {
	data, err := json.MarshalIndent(expense, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile("expenses.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadBudgets() ([]model.Budget, error) {
	_, err := os.Stat("budget.json")
	if os.IsNotExist(err) {
		return []model.Budget{}, nil
	}

	file, err := os.ReadFile("budget.json")
	if err != nil {
		return []model.Budget{}, err
	}

	var budget []model.Budget

	err = json.Unmarshal(file, &budget)
	if err != nil {
		return []model.Budget{}, err
	}

	return budget, nil
}

func SaveBudgets(budget []model.Budget) error {
	data, err := json.MarshalIndent(budget, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile("budget.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}