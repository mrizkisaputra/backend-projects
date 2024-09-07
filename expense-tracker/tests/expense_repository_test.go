package tests

import (
	"expense-tracker/cmd/entities"
	"expense-tracker/cmd/repositories"
	"expense-tracker/internal/log"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var expenseRepo = repositories.NewExpenseRepository(log.Log)

func TestAdd(t *testing.T) {
	date, _ := time.Parse("2006-01-02", time.Now().Local().Format("2006-01-02"))
	s := []entities.Expense{
		{Id: 3, Description: "Launch3", Amount: 30.00, Category: "food", Date: date},
		{Id: 2, Description: "Launch2", Amount: 20.00, Category: "food", Date: date},
		{Id: 1, Description: "Launch1", Amount: 10.00, Category: "food", Date: date},
	}

	for i, e := range s {
		t.Run(fmt.Sprintf("Sub Test-%d", i), func(t *testing.T) {
			status, err := expenseRepo.Add(e)
			assert.Nil(t, err)
			assert.True(t, status)
		})
	}
}

func TestGetAll(t *testing.T) {
	expenses, err := expenseRepo.GetAll()
	assert.Nil(t, err)
	assert.NotNil(t, expenses)
	fmt.Println(expenses)
}

func TestGetSummary(t *testing.T) {
	totalAmount, err := expenseRepo.GetSummary()
	assert.Nil(t, err)
	fmt.Println(totalAmount)
}

func TestGetSummaryByMonthCurrentYear(t *testing.T) {
	s := []struct {
		Month string
	}{
		{Month: "08"},
		{Month: "09"},
		{Month: "10"},
	}

	for i, v := range s {
		t.Run(fmt.Sprintf("Sub Test-%d", i), func(t *testing.T) {
			totalAmount, err := expenseRepo.GetSummaryByMonthCurrentYear(v.Month)
			assert.Nil(t, err)
			fmt.Println(totalAmount)
		})
	}
}

func TestDelete(t *testing.T) {
	s := []struct {
		Id      int
		IsExist bool
	}{
		{Id: 2, IsExist: true},
		{Id: 100, IsExist: false},
	}

	for i, v := range s {
		t.Run(fmt.Sprintf("Sub Test-%d", i), func(t *testing.T) {
			success, err := expenseRepo.Delete(v.Id)
			assert.Nil(t, err)
			assert.Equal(t, v.IsExist, success)
		})
	}
}

func TestUpdate(t *testing.T) {
	s := []struct {
		Id      int
		IsExist bool
	}{
		{Id: 2, IsExist: true},
		{Id: 100, IsExist: false},
	}

	for i, v := range s {
		t.Run(fmt.Sprintf("Sub Test-%d", i), func(t *testing.T) {
			success, err := expenseRepo.Update(v.Id, "", 200)
			assert.Nil(t, err)
			assert.Equal(t, v.IsExist, success)
		})
	}
}

func TestGetByCategory(t *testing.T) {
	categories, err := expenseRepo.GetByCategory("food")
	assert.Nil(t, err)
	assert.NotNil(t, categories)
	fmt.Println(categories)
}
