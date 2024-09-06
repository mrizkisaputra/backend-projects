package example

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestExample1(t *testing.T) {
	wd, _ := os.Getwd()
	join := filepath.Join(wd, "expense.json")
	fmt.Println(join)
}
