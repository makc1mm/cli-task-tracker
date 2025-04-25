package domain

type Operation string

const (
	OperationAdd    Operation = "add"
	OperationUpdate Operation = "update"
	OperationDelete Operation = "delete"
	OperationMark   Operation = "mark"
	OperationList   Operation = "list"
)
