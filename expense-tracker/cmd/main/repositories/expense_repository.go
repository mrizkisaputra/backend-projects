package repositories

import (
	. "expense-tracker/cmd/main/entities"
	"expense-tracker/internal/exceptions"
	"expense-tracker/internal/filesystem"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

var instance ExpenseRepositoryInterface

type expenseRepository struct {
	log *logrus.Logger
}

func NewExpenseRepository(log *logrus.Logger) ExpenseRepositoryInterface {
	if instance == nil {
		instance = &expenseRepository{
			log: log,
		}
	}
	return instance
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

	if len(data) == 0 {
		return nil, exceptions.NewErrNotFound("Data expense empty!")
	}

	return data, nil
}

func (e *expenseRepository) GetSummary() (float64, error) {
	/* read file */
	data, errRead := filesystem.ReadFile()
	if errRead != nil {
		return 0, errRead
	}

	if len(data) == 0 {
		return 0, exceptions.NewErrNotFound("Data expense empty!")
	}

	/* ambil amount expense */
	totalAmount := 0.0
	for _, item := range data {
		amount := item.Amount
		/* jumlahkan amount dari tiap expense */
		totalAmount += amount
	}
	return totalAmount, nil
}

func (e *expenseRepository) GetSummaryByMonthCurrentYear(month string) (float64, error) {
	/* read file */
	data, errRead := filesystem.ReadFile()
	if errRead != nil {
		return 0, errRead
	}

	if len(data) == 0 {
		return 0, exceptions.NewErrNotFound("Data expense empty!")
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

	return totalAmount, nil
}

func (e *expenseRepository) Delete(id int) (bool, error) {
	/* read file */
	data, errRead := filesystem.ReadFile()
	if errRead != nil {
		return false, errRead
	}

	if len(data) == 0 {
		return false, exceptions.NewErrNotFound("Data expense empty!")
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
		return isExist, exceptions.NewErrNotFound(fmt.Sprintf("Data with ID %d not found!", id))
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

	if len(data) == 0 {
		return false, exceptions.NewErrNotFound("Data expense empty!")
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
		return isExist, exceptions.NewErrNotFound(fmt.Sprintf("Data with ID %d not found!", id))
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

	if len(data) == 0 {
		return nil, exceptions.NewErrNotFound("Data expense empty!")
	}

	/* cari data expense berdasarkan parameter category */
	var listCategories []Expense
	var exist bool
	for _, item := range data {
		if item.Category == category {
			exist = true
			listCategories = append(listCategories, item)
		}
	}

	if !exist {
		return nil, exceptions.NewErrNotFound(fmt.Sprintf("list of category %s not found", category))
	}

	/* return list category */
	return listCategories, nil
}
