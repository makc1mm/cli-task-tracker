package operations

import (
	"encoding/json"
	"github.com/makc1mm/cli-task-tracker/internal/domain"
	"io"
	"os"
)

func getTasksFromFile(file *os.File) ([]domain.Task, error) {
	decoder := json.NewDecoder(file)
	tasks := []domain.Task{}
	if err := decoder.Decode(&tasks); err != nil && err != io.EOF {
		return nil, err
	}
	return tasks, nil
}

func rewriteTasksToFile(file *os.File, tasks []domain.Task) error {
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}
	if err := file.Truncate(0); err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tasks); err != nil {
		return err
	}

	return nil
}
