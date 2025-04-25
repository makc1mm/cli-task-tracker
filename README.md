# First golang project. Just syntax learning

## How to use

1. In root project directory run `go build -o task-cli .\cmd\main.go`
2. There are some operations:
    - Adding a new task: `task-cli add "Buy groceries"`
    - Updating tasks: `task-cli update 1 "Buy groceries and cook dinner"`
    - Deleting tasks: `task-cli delete 1`
    - Marking a task as in-progress or done: `task-cli mark 1 in-progress`
    - Listing all tasks: `task-cli list`
    - Listing tasks by status: `task-cli list todo`

P.S. there are some bugs, will fix it later :)