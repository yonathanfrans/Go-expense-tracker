package main

import (
	"example/expense-tracker/service"
	"fmt"
	"os"
	"strconv"
)

func getFlagValue(args []string, flag string) string {
	var flagValue string

	for i := 2; i < len(args); i++ {
		if args[i] == flag && i+1 < len(args) {
			flagValue = args[i+1]
		}
	}

	return flagValue
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Silahkan masukkan perintah! Contoh: add \"Belajar Go\" atau list")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 8 {
			fmt.Println("Error: Argumen kurang! Contoh: add --description \"Makan\" --amount 50000")
			return
		}

		description := getFlagValue(os.Args, "--description")
		amountStr := getFlagValue(os.Args, "--amount")
		category := getFlagValue(os.Args, "--category")

		if description == "" || amountStr == "" || category == "" {
			fmt.Println("Error: Flag --description, --amount dan --category wajib diisi!")
			return
		}

		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			fmt.Println("Error: Nominal amount harus berupa angka!")
			return
		}

		newID, err := service.AddExpense(description, amount, category)
		if err != nil {
			fmt.Printf("Error: Gagal menambah expense: %v\n", err)
			return
		}

		fmt.Printf("Expense added successfully (ID: %d)\n", newID)

	case "update":
		if len(os.Args) < 10 {
			fmt.Println("Error: ID, deskripsi, amount dan category tidak boleh kosong! Contoh: update --id 1 --description \"Sarapan\" --amount 30000 --category ")
			return
		}

		idStr := getFlagValue(os.Args, "--id")
		if idStr == "" {
			fmt.Println("Error: Flag --id tidak ditemukan id")
			return
		}

		description := getFlagValue(os.Args, "--description")
		category := getFlagValue(os.Args, "--category")
		amountStr := getFlagValue(os.Args, "--amount")
		if description == "" || amountStr == "" || category == "" {
			fmt.Println("Error: Flag --description, --amount, dan --category wajib diisi!")
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("ID harus integer!")
			return
		}

		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			fmt.Println("Error: Nominal amount harus berupa angka!")
			return
		}

		err = service.UpdateExpense(id, description, amount, category)
		if err != nil {
			fmt.Printf("Gagal memperbarui expense: %v\n", err)
		}

		fmt.Printf("Expense updated successfully (ID: %d)\n", id)

	case "list":
		categoryFilter := getFlagValue(os.Args, "--category")

		err := service.ListExpenses(categoryFilter)
		if err != nil {
			fmt.Printf("Gagal menampilkan expense: %v\n", err)
		}
		
	case "delete":
		if len(os.Args) < 4 {
			fmt.Println("Error: Argumen kurang! Contoh: delete --id 1")
			return
		}

		idStr := getFlagValue(os.Args, "--id")
		if idStr == "" {
			fmt.Println("Error: Flag --id tidak ditemukan id")
			return
		}

		id, err := strconv.Atoi(idStr) 
		if err != nil {
			fmt.Println("ID harus integer!")
			return
		}

		err = service.DeleteExpense(id)
		if err != nil {
			fmt.Printf("Gagal menghapus expense: %v\n", err)
			return
		}

		fmt.Println("Expense deleted successfully")

	case "summary":
		monthFlag := getFlagValue(os.Args, "--month")
		month := 0

		var err error
		if monthFlag != "" {
			month, err = strconv.Atoi(monthFlag)
			if err != nil || month < 1 || month > 12 {
				fmt.Println("Bulan harus berupa angka antara 1 - 12!")
				return
			}

		}

		total, err := service.SummaryExpenses(month)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if month == 0 {
			fmt.Printf("Total expenses: Rp.%.2f\n", total)
		} else {
			fmt.Printf("Total expenses for month %d: Rp.%.2f\n", month, total)
		}

	case "set-budget":
		monthFlag := getFlagValue(os.Args, "--month")
		amountStr := getFlagValue(os.Args, "--amount")
		
		if monthFlag == "" || amountStr == "" {
			fmt.Println("Error: Flag --month dan --amount wajib diisi!")
			return
		}

		month, err := strconv.Atoi(monthFlag)
		if err != nil || month < 1 || month > 12 {
			fmt.Println("Bulan harus berupa angka antara 1 - 12!")
			return
		}

		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			fmt.Println("Error: Nominal amount harus berupa angka!")
			return
		}

		err = service.SetMonthlyBudget(month, amount)
		if err != nil {
			fmt.Printf("Gagal set budget bulan ini: %v\n", err)
		}

		fmt.Println("Set budget this month successfully")

	case "export":
		err := service.ExportToCSV()
		if err != nil {
			fmt.Printf("Gagal melakukan export CSV: %v\n", err)
			return
		}

		fmt.Println("Expenses exported successfully to \"expenses.csv\"")
	}
}