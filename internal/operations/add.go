package operations

import (
	"github.com/makc1mm/cli-task-tracker/internal/domain"
	"os"
	"time"
)

func Add(filePath string, description string) (int, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE, 0666)
	if err != nil {
		return 0, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("Error occurred while closing file")
		}
	}(file)

	tasks, err := getTasksFromFile(file)
	if err != nil {
		return 0, err
	}

	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	newId := maxID + 1
	newTask := domain.Task{
		ID:          newId,
		Description: description,
		Status:      domain.StatusTodo,
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   time.Now().Local(),
	}

	tasks = append(tasks, newTask)

	if err := rewriteTasksToFile(file, tasks); err != nil {
		return 0, err
	}

	return newId, nil
}
