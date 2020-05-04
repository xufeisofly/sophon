package sophon

type Task struct {
	Desc string
	Fn   BlockFunc
}

func NewTask(desc string, fn BlockFunc) *Task {
	return &Task{desc, fn}
}
