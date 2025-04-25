package main

import (
	"encoding/json"
	"fmt"
	"github.com/makc1mm/cli-task-tracker/internal/domain"
	"github.com/makc1mm/cli-task-tracker/internal/operations"
	"os"
	"strconv"
)

/**
операции:
	add
	update
	delete
	mark
	list
*/

const (
	filePath string = "tasks.json"
)

func main() {
	switch os.Args[1] {
	case string(domain.OperationAdd):
		if len(os.Args) < 3 {
			fmt.Printf("Usage: task-cli %s \"task description\"\n", string(domain.OperationAdd))
			return
		}

		description := os.Args[2]

		id, err := operations.Add(filePath, description)
		if err != nil {
			handleError(domain.OperationAdd, err)
			return
		}
		fmt.Printf("Task added successfully (ID: %d)\n", id)

	case string(domain.OperationUpdate):
		if len(os.Args) < 4 {
			fmt.Printf("Usage: task-cli %s <id> \"task description\"\n", string(domain.OperationUpdate))
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			handleError(domain.OperationUpdate, err)
		}

		description := os.Args[3]

		if err := operations.Update(filePath, id, description); err != nil {
			handleError(domain.OperationUpdate, err)
			return
		}

	case string(domain.OperationDelete):
		if len(os.Args) < 3 {
			fmt.Printf("Usage: task-cli %s <id>\n", string(domain.OperationDelete))
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			handleError(domain.OperationDelete, err)
		}

		if err := operations.Delete(filePath, id); err != nil {
			handleError(domain.OperationDelete, err)
			return
		}

	case string(domain.OperationList):
		status := domain.StatusAll
		if len(os.Args) > 2 {
			status = domain.Status(os.Args[2])
		}

		tasks, err := operations.List(filePath, status)
		if err != nil {
			handleError(domain.OperationDelete, err)
			return
		}

		for _, t := range tasks {
			jsonTask, err := json.MarshalIndent(t, "", "  ")
			if err != nil {
				handleError(domain.OperationDelete, err)
				return
			}
			fmt.Printf("%v\n", string(jsonTask))
		}

	case string(domain.OperationMark):
		if len(os.Args) < 4 {
			fmt.Printf("Usage: task-cli %s <id> done\n", string(domain.OperationMark))
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			handleError(domain.OperationMark, err)
		}

		status := domain.Status(os.Args[3])

		if err := operations.Mark(filePath, id, status); err != nil {
			handleError(domain.OperationMark, err)
			return
		}

	default:
		os.Exit(1)
	}

}

func handleError(operation domain.Operation, err error) {
	fmt.Printf("Error occurred while \"%s\" operation: %v\n", string(operation), err)
}
