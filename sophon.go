package sophon

import (
	"fmt"
	"os"
)

type Sophon struct {
	Tasks []*Task
}

// New 实例化 sophon
func New() *Sophon {
	return &Sophon{}
}

type BlockFunc func() bool

// Block 注册监测任务
func (s *Sophon) Block(desc string, f BlockFunc) {
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

type result struct {
	desc    string
	canPass bool
}

func (s *Sophon) Run() {
	var rets = make(chan result, len(s.Tasks))
	var errText string

	for _, task := range s.Tasks {
		go func(t *Task) {
			canPass := t.Fn()
			rets <- result{
				desc:    t.Desc,
				canPass: canPass,
			}
		}(task)
	}

	for range s.Tasks {
		ret := <-rets
		if !ret.canPass {
			errText = errText + fmt.Sprintf("Sophon blocked!: %s", ret.desc) + "\n"
		}
	}

	if errText != "" {
		panic(errText)
	}
}

func checkExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(err)
	}
}
