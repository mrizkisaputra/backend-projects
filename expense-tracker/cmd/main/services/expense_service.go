package services

import (
	. "expense-tracker/cmd/main/entities"
	. "expense-tracker/cmd/main/repositories"
	"expense-tracker/internal/exceptions"
	"expense-tracker/internal/filesystem"
	"expense-tracker/internal/log"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
	"time"
)

var instance ExpenseServiceInterface

type expenseService struct {
	Repository ExpenseRepositoryInterface
	validate   *validator.Validate
}

func NewExpenseService(expenseRepo ExpenseRepositoryInterface) ExpenseServiceInterface {
	if instance == nil {
		instance = &expenseService{
			Repository: expenseRepo,
			validate:   validator.New(),
		}
	}
	return instance
}

func (e *expenseService) Add(description string, amount float64, category string) (string, error) {
	/* validasi data parameter */
	if err := e.validate.Var(description, "required"); err != nil {
		return "", exceptions.NewErrValidation("description is required")
	}

	if err := e.validate.Var(amount, "gt=0"); err != nil {
		return "", exceptions.NewErrValidation("amount must be greater than zero")
	}

	if err := e.validate.Var(category, "alpha"); err != nil {
		return "", exceptions.NewErrValidation("category must be alphabet")
	}

	/* siapkan data expense date */
	dateFormated := time.Now().Local().Format("2006-01-02")
	parsedTime, errParsedTime := time.Parse("2006-01-02", dateFormated)
	if errParsedTime != nil {
		log.Log.Errorf("Error parsing date: %v", errParsedTime.Error())
		return "", errParsedTime
	}

	/* siapkan data expense id */
	var expenseId int
	data, errRead := filesystem.ReadFile()
	if errRead != nil {
		return "", errRead
	}

	if len(data) == 0 {
		expenseId = 1
	} else {
		increment := data[len(data)-1].Id + 1
		expenseId = increment
	}
	expense := Expense{
		Id:          expenseId,
		Description: description,
		Amount:      amount,
		Category:    category,
		Date:        parsedTime,
	}
	/* lanjutkan ke repository untuk disimpan */
	success, err := e.Repository.Add(expense)
	if err != nil {
		return "", err
	}

	if !success {
		return fmt.Sprintf("Expense added failed (ID: %d)", expenseId), nil
	}
	return fmt.Sprintf("Expense added successfully (ID: %d)", expenseId), nil
}

func (e *expenseService) GetAll() (string, error) {
	data, err := e.Repository.GetAll()
	if err != nil {
		return "", err
	}

	var strBuilder = new(strings.Builder)
	strBuilder.WriteString(`
ID              Date               Description     Amount    Category
`)
	for _, item := range data {
		str := fmt.Sprintf(`
%d  %v  %s      $%.2f  %s
`, item.Id, item.Date, item.Description, item.Amount, item.Category)
		strBuilder.WriteString(str)
	}
	return strBuilder.String(), nil
}

func (e *expenseService) GetSummary() (string, error) {
	summaryTotalAmount, err := e.Repository.GetSummary()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Total expenses: $%.2f", summaryTotalAmount), nil
}

func (e *expenseService) GetSummaryByMonthCurrentYear(month string) (string, error) {
	/* validasi data parameter */
	if err := e.validate.Var(month, "numeric"); err != nil {
		return "", exceptions.NewErrValidation("Month must be numeric and start '0'")
	}

	// Cek apakah bulan dari 01 sampai 09, atau 10, 11, 12
	validMonths := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	isValidMonth := false

	for _, validMonth := range validMonths {
		if month == validMonth {
			isValidMonth = true
			break
		}
	}

	if !isValidMonth {
		return "", exceptions.NewErrValidation(`Invalid month format. 
Please input in the format'01 to 12'. january - september with prefix '0'`)
	}

	summaryTotalAmount, err := e.Repository.GetSummaryByMonthCurrentYear(month)
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(month, "0") {
		month, _ = strings.CutPrefix(month, "0")
	}
	parseInt, _ := strconv.ParseInt(month, 10, 64)
	return fmt.Sprintf("Total expenses for %v: $%.2f", time.Month(parseInt), summaryTotalAmount), nil
}

func (e *expenseService) Delete(id int) (string, error) {
	/* validasi data parameter */
	if err := e.validate.Var(id, "gt=0"); err != nil {
		return "", exceptions.NewErrValidation("Id must be grater then 0")
	}

	success, errDelete := e.Repository.Delete(id)
	if errDelete != nil {
		return "", errDelete
	}

	if !success {
		return fmt.Sprintf("Expense deleted failed (ID: %d)", id), nil
	}
	return fmt.Sprintf("Expense deleted successfully (ID: %d)", id), nil
}

func (e *expenseService) Update(id int, description string, amount float64) (string, error) {
	/* validasi data parameter */
	if err := e.validate.Var(id, "gt=0"); err != nil {
		return "", exceptions.NewErrValidation("Please input correct id, must be grater then 0")
	}

	if err := e.validate.Var(amount, "gt=0"); err != nil {
		return "", exceptions.NewErrValidation("amount must be greater than zero")
	}

	success, errUpdate := e.Repository.Update(id, description, amount)
	if errUpdate != nil {
		return "", errUpdate
	}

	if !success {
		return fmt.Sprintf("Expense updated failed (ID: %d)", id), nil
	}
	return fmt.Sprintf("Expense updated successfully (ID: %d)", id), nil
}

func (e *expenseService) GetByCategory(category string) (string, error) {
	if err := e.validate.Var(category, "alpha"); err != nil {
		return "", exceptions.NewErrValidation("category must be alphabet")
	}

	data, err := e.Repository.GetByCategory(category)
	if err != nil {
		return "", err
	}

	var strBuilder = new(strings.Builder)
	for _, item := range data {
		str := fmt.Sprintf(`
%d | %v | %s      | $%.2f | %s
`, item.Id, item.Date, item.Description, item.Amount, item.Category)
		strBuilder.WriteString(str)
	}
	return strBuilder.String(), nil
}
