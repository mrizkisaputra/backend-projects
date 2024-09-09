package repositories

import (
	. "expense-tracker/cmd/main/entities"
)

type ExpenseRepositoryInterface interface {
	Add(expense Expense) (bool, error)
	GetAll() ([]Expense, error)
	GetSummary() (float64, error)
	GetSummaryByMonthCurrentYear(month string) (float64, error)
	Delete(id int) (bool, error)
	Update(id int, newDescription string, newAmount float64) (bool, error)
	GetByCategory(category string) ([]Expense, error)
}
