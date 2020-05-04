package sophon

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type File struct {
	path string
	info os.FileInfo
}

func NewFile(path string, info os.FileInfo) *File {
	checkExist(path)
	info, err := os.Lstat(path)
	if err != nil {
		panic(err)
	}
	return &File{path, info}
}

func (f *File) Name() string {
	return f.info.Name()
}

func (f *File) Ext() string {
	return filepath.Ext(f.path)
}

func (f *File) String() string {
	return f.path
}

func (f *File) Content() string {
	b, err := ioutil.ReadFile(f.path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
