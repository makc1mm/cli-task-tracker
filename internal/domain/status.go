package domain

type Status string

const (
	StatusAll        Status = "all"
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)
