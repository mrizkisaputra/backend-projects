package services

type ExpenseServiceInterface interface {
	Add(description string, amount float64, category string) (string, error)
	GetAll() (string, error)
	GetSummary() (string, error)
	GetSummaryByMonthCurrentYear(month string) (string, error)
	Delete(id int) (string, error)
	Update(id int, description string, amount float64) (string, error)
	GetByCategory(category string) (string, error)
}
