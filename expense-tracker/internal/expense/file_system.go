package expense

import (
	"encoding/json"
	"expense-tracker/cmd/entities"
	. "expense-tracker/internal/log"
	"os"
	"path/filepath"
)

var filename = "expenses.json"

// getFilePath return current working directory,
// the directory location where the file was created.
func getFilePath() string {
	if workingDirectory, err := os.Getwd(); err != nil {
		Log.Errorf("Error getting current working directory: %v", err.Error())
		return ""
	} else {
		/* mengabungkan antara jalur file & nama file*/
		return filepath.Join(workingDirectory, filename)
	}
}

func ReadFile() ([]entities.Expense, error) {
	pathLocation := getFilePath()

	/* detect whether the file already exists */
	_, err := os.Stat(pathLocation)

	/* if file does not exist, create file*/
	if os.IsNotExist(err) {
		file, err := os.Create(pathLocation)
		if err != nil {
			Log.Errorf("Error creating file: %v", err.Error())
			return nil, err
		}
		defer file.Close()

		/* write the contents of the file for the first time */
		errWriteFile := os.WriteFile(pathLocation, []byte("[]"), os.ModeAppend)
		if errWriteFile != nil {
			Log.Errorf("Error write contents of the file for the first time: %v", errWriteFile.Error())
			return nil, errWriteFile
		}
		return []entities.Expense{}, nil
	}

	/* file already exist */
	fileOpened, errOpened := openFile(pathLocation)
	defer fileOpened.Close()
	if errOpened != nil {
		Log.Errorf("Error opening file: %v", errOpened.Error())
		return nil, errOpened
	}

	/* decode contents file to slice */
	var data []entities.Expense
	errDecoded := json.NewDecoder(fileOpened).Decode(&data)
	if errDecoded != nil {
		Log.Errorf("Error decoding file: %v", errDecoded.Error())
		return nil, errDecoded
	}
	return data, nil
}

func openFile(pathLocation string) (*os.File, error) {
	fileOpened, errOpenedFile := os.OpenFile(pathLocation, os.O_RDWR, 0644)
	if errOpenedFile != nil {
		return nil, errOpenedFile
	}
	return fileOpened, nil
}

func WriteFile(data []entities.Expense) (bool, error) {
	pathLocation := getFilePath()

	/* open file */
	fileOpened, errOpened := openFile(pathLocation)
	defer fileOpened.Close()
	if errOpened != nil {
		Log.Errorf("Error opening file: %v", errOpened.Error())
		return false, errOpened
	}

	/* encode data in file content */
	errEncoded := json.NewEncoder(fileOpened).Encode(data)
	if errEncoded != nil {
		Log.Errorf("Error encoding data to file: %v", errEncoded.Error())
		return false, errEncoded
	}

	return true, nil
}
