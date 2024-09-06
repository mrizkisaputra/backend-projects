package tests

import (
	"expense-tracker/cmd/entities"
	"expense-tracker/internal/expense"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestReadFile(t *testing.T) {
	data, err := expense.ReadFile()
	assert.Nil(t, err)
	assert.NotNil(t, data)
}

func TestWriteFile(t *testing.T) {
	data := []entities.Expense{
		{Id: 1, Description: "Testing write file", Amount: 10.10, Date: time.Now(), Category: "test"},
		{Id: 2, Description: "Testing write file", Amount: 20.10, Date: time.Now(), Category: "test"},
		{Id: 3, Description: "Testing write file", Amount: 40.10, Date: time.Now(), Category: "test"},
	}

	status, err := expense.WriteFile(data)
	assert.Nil(t, err)
	assert.True(t, status)
}
