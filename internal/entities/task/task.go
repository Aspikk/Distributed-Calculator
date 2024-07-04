package task

var (
	id = 0
)

type Task struct {
	ID             int         `json:"id"`
	Arg1           interface{} `json:"arg1"`
	Arg2           interface{} `json:"arg2"`
	Operation      string      `json:"operation"`
	Operation_time string      `json:"operation_time"`
}

func New(arg1, arg2 interface{}, operation, operation_time string) *Task {
	id++
	return &Task{
		ID:             id,
		Arg1:           arg1,
		Arg2:           arg2,
		Operation:      operation,
		Operation_time: operation_time,
	}
}

func IsTask(entity interface{}) bool {
	_, ok := entity.(*Task)
	return ok
}
