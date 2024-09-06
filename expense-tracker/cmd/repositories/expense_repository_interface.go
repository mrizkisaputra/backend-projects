package repositories

import . "expense-tracker/cmd/entities"

type ExpenseRepositoryInterface interface {
	Add(expense Expense) (bool, error)
	GetAll() ([]string, error)
	GetSummary() (float64, error)
	GetSummaryByMonth(month string) (float64, error)
	Delete(id int) (bool, error)
	Update(id int, description string, amount float64) (bool, error)
	GetByCategory(category string) ([]Expense, error)
}
