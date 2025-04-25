package operations

import (
	"github.com/makc1mm/cli-task-tracker/internal/domain"
	"os"
)

func List(filePath string, status domain.Status) ([]domain.Task, error) {
	tasks, err := listByStatus(filePath, status)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func listByStatus(filePath string, status domain.Status) ([]domain.Task, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("Error occurred while closing file")
		}
	}(file)

	tasks, err := getTasksFromFile(file)
	if err != nil {
		return nil, err
	}

	newTasks := []domain.Task{}
	if status != domain.StatusAll {
		for _, t := range tasks {
			if t.Status == status {
				newTasks = append(newTasks, t)
			}
		}
	} else {
		return tasks, nil
	}

	return newTasks, nil
}
