package sophon

import (
	"os"
)

type Sophon struct {
	Tasks []*Task
}

// New 实例化 sophon
func New() *Sophon {
	return &Sophon{}
}

// Block 注册监测任务
func (s *Sophon) Block(desc string, f func()) {
	task := NewTask(desc, f)
	s.register(task)
}

func (s *Sophon) register(task *Task) {
	s.Tasks = append(s.Tasks, task)
}

func (s *Sophon) Dir(path string) *Dir {
	return NewDir(path, nil)
}

func (s *Sophon) File(path string) *File {
	return NewFile(path, nil)
}

func (s *Sophon) Run() {
	for _, task := range s.Tasks {
		task.Fn()
	}
}

func checkExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(err)
	}
}
