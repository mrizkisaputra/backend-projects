package example

import (
	"expense-tracker/internal/filesystem"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestExample1(t *testing.T) {
	wd, _ := os.Getwd()
	join := filepath.Join(wd, "filesystem.json")
	fmt.Println(join)
}

func TestExample2(t *testing.T) {
	date, _ := time.Parse("2006-01-02", time.Now().Local().Format("2006-01-02"))
	dateStr := date.Format("2006-01")
	fmt.Println(dateStr)
}

func TestExample3(t *testing.T) {
	data, errRead := filesystem.ReadFile()
	if errRead != nil {
		fmt.Println(errRead)
	}

	for i, item := range data {
		if item.Id == 2 {
			data = append(data[:i], data[i+1:]...)
			fmt.Printf(fmt.Sprintf("Expense deleted: %s, Amount: %.2f, Category: %s (ID: %d)\n", item.Description, item.Amount, item.Category, 2))
			break
		}
	}
	fmt.Println(data)

	//_, errWrite := filesystem.WriteFile(deleted)
	//if errWrite != nil {
	//	fmt.Println(errWrite)
	//}
}

func TestExample4(t *testing.T) {
	arr := []int{10, 20, 30, 40, 50}

	arr = append(arr[:2], arr[3:]...)
	fmt.Println(arr)
}
