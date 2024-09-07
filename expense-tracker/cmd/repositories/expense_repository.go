package repositories

import (
	"expense-tracker/internal/filesystem"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)
import . "expense-tracker/cmd/entities"

type expenseRepository struct {
	log *logrus.Logger
}

func NewExpenseRepository(log *logrus.Logger) ExpenseRepositoryInterface {
	return &expenseRepository{
		log: log,
	}
}

func (e *expenseRepository) Add(expense Expense) (bool, error) {
	/* ambil data yang sudah ada / data kosong, dengan cara membaca file */
	data, errRead := filesystem.ReadFile()
	if errRead != nil {
		return false, errRead
	}

	/* tambahkan data baru, setelah data yang sudah ada / belum ada */
	data = append(data, expense)

	/* tulis kembali file, untuk menambahkan data baru */
	status, errWrite := filesystem.WriteFile(data)
	if errWrite != nil {
		return status, errWrite
	}
	return status, nil
}

func (e *expenseRepository) GetAll() ([]Expense, error) {
	data, errRead := filesystem.ReadFile()
	if errRead != nil {
		return nil, errRead
	}
	return data, nil
}

func (e *expenseRepository) GetSummary() (float64, error) {
	/* read file */
	data, errRead := filesystem.ReadFile()
	if errRead != nil {
		return 0, errRead
	}

	/* ambil amount expense */
	totalAmount := 0.0
	for _, item := range data {
		amount := item.Amount
		/* jumlahkan amount dari tiap expense */
		totalAmount += amount
	}

	totalAmountStr := fmt.Sprintf("%.2f", totalAmount)
	totalAmountParsed, _ := strconv.ParseFloat(totalAmountStr, 64)
	return totalAmountParsed, nil
}

func (e *expenseRepository) GetSummaryByMonthCurrentYear(month string) (float64, error) {
	/* read file */
	data, errRead := filesystem.ReadFile()
	if errRead != nil {
		return 0, errRead
	}

	/* cari berdasarkan parameter month */
	totalAmount := 0.0
	for _, item := range data {
		dateStr := item.Date.Format("2006-01")
		currentYear := time.Now().Local().Year()
		monthCurrentYear := fmt.Sprintf("%d-%s", currentYear, month)

		/* ambil amount expense */
		if monthCurrentYear == dateStr {
			amount := item.Amount
			/* jumlahkan amount  */
			totalAmount += amount
		}
	}

	totalAmountStr := fmt.Sprintf("%.2f", totalAmount)
	totalAmountParsed, _ := strconv.ParseFloat(totalAmountStr, 64)
	return totalAmountParsed, nil
}

func (e *expenseRepository) Delete(id int) (bool, error) {
	/* read file */
	data, errRead := filesystem.ReadFile()
	if errRead != nil {
		return false, errRead
	}

	/* cari data expense berdasarkan parameter id */
	isExist := false
	for i, item := range data {
		if item.Id == id {
			isExist = true
			/* hapus data expense */
			data = append(data[:i], data[i+1:]...)
			break
		}
	}

	if !isExist {
		return isExist, nil
	}

	/* setelah dihapus, tulis ulang */
	status, errWrite := filesystem.WriteFile(data)
	if errWrite != nil {
		return status, errWrite
	}
	return true, nil
}

func (e *expenseRepository) Update(id int, newDescription string, newAmount float64) (bool, error) {
	/* read file */
	data, errRead := filesystem.ReadFile()
	if errRead != nil {
		return false, errRead
	}

	/* cari data expense berdasarkan id expense yang ingin diubah */
	var isExist bool
	for i, item := range data {
		if item.Id == id {
			isExist = true
			var actionDesc string
			var actionAmount float64
			if newDescription == "" {
				actionDesc = data[i].Description
			} else {
				actionDesc = newDescription
			}

			if newAmount == 0 {
				actionAmount = data[i].Amount
			} else {
				actionAmount = newAmount
			}

			newExpense := Expense{
				Id:          data[i].Id,
				Description: actionDesc,
				Amount:      actionAmount,
				Date:        data[i].Date,
				Category:    data[i].Category,
			}

			/* update data */
			data[i] = newExpense
			break
		}
	}

	if !isExist {
		return isExist, nil
	}

	/* seletelah diubah, tulis ulang */
	status, errWrite := filesystem.WriteFile(data)
	if errWrite != nil {
		return status, errWrite
	}
	return isExist, nil
}

func (e *expenseRepository) GetByCategory(category string) ([]Expense, error) {
	/* read file */
	data, errRead := filesystem.ReadFile()
	if errRead != nil {
		return nil, errRead
	}

	/* cari data expense berdasarkan parameter category */
	var listCategories []Expense
	for _, item := range data {
		if item.Category == category {
			listCategories = append(listCategories, item)
		}
	}

	/* return list category */
	return listCategories, nil
}
