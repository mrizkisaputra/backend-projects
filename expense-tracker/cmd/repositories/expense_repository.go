package repositories

import "github.com/sirupsen/logrus"
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
	//TODO implement me
	panic("implement me")
}

func (e *expenseRepository) GetAll() ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (e *expenseRepository) GetSummary() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (e *expenseRepository) GetSummaryByMonth(month string) (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (e *expenseRepository) Delete(id int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (e *expenseRepository) Update(id int, description string, amount float64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (e *expenseRepository) GetByCategory(category string) ([]Expense, error) {
	//TODO implement me
	panic("implement me")
}
