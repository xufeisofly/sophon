package sophon_test

import (
	"testing"

	"github.com/xufeisofly/sophon"
)

func TestDir(t *testing.T) {
	s := sophon.New()
	dir := s.Dir("./demo")

	if len(dir.AllFiles()) != 3 {
		t.Errorf("all files number expected: %d, actual: %d", 3, len(dir.AllFiles()))
	}

	if len(dir.Files()) != 1 {
		t.Errorf("files number expected: %d, actual: %d", 1, len(dir.Files()))
	}

	if len(dir.Dirs()) != 1 {
		t.Errorf("dirs number expected: %d, actual: %d", 1, len(dir.Dirs()))
	}

	if len(dir.AllDirs()) != 2 {
		t.Errorf("all dirs number expected: %d, actual: %d", 2, len(dir.AllDirs()))
	}
}

func TestFile(t *testing.T) {
	s := sophon.New()
	file := s.File("./demo/pkg/foo_controller.go")

	if file.Ext() != ".go" {
		t.Errorf("file ext expected: %s, actual: %s", ".go", file.Ext())
	}
}

func TestRun(t *testing.T) {
	assertPanic(t, func() {
		s := sophon.New()

		s.Block("test1", func() bool {
			return false
		})

		s.Block("test2", func() bool {
			return false
		})

		if len(s.Tasks) != 2 {
			t.Errorf("[TestBlock] tasks number expected: %d, actual: %d", 1, len(s.Tasks))
		}

		s.Run()
	})
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
