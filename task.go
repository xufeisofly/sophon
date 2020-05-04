package sophon

type Task struct {
	Desc string
	Fn   func()
}

func NewTask(desc string, fn func()) *Task {
	return &Task{desc, fn}
}
