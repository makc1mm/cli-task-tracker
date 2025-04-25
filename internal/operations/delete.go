package operations

import "os"

func Delete(filePath string, id int) error {
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

	tasks, err := getTasksFromFile(file)
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	if err := rewriteTasksToFile(file, tasks); err != nil {
		return err
	}

	return nil
}
