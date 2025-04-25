package operations

import (
	"encoding/json"
	"github.com/makc1mm/cli-task-tracker/internal/domain"
	"io"
	"os"
	"time"
)

func Mark(filePath string, id int, status domain.Status) error {
	if status == domain.StatusAll {
		panic("Use another status")
	}

	file, err := os.OpenFile(filePath, os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("Error occurred while closing file")
		}
	}(file)

	tasks := []domain.Task{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil && err != io.EOF {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now().Local()
			break
		}
	}

	if err := rewriteTasksToFile(file, tasks); err != nil {
		return err
	}

	return nil
}
