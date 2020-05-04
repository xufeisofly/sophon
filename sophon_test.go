package sophon_test

import (
	"fmt"
	"sophon"
	"testing"
)

func TestBlock(t *testing.T) {
	s := sophon.New()
	s.Block("test1", func() {
		fmt.Println("show test1")
	})

	if len(s.Tasks) != 1 {
		t.Errorf("[TestBlock] tasks number expected: %d, actual: %d", 1, len(s.Tasks))
	}
}

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
	fmt.Println(file.Content())
}
