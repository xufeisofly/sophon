package sophon

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type Dir struct {
	path     string
	files    []*File
	allFiles []*File
	dirs     []*Dir
	allDirs  []*Dir

	info os.FileInfo
}

func NewDir(path string, info os.FileInfo) *Dir {
	checkExist(path)
	files, dirs := walkFiles(path)
	allFiles, allDirs := walkAllFiles(path)

	return &Dir{
		path:     path,
		files:    files,
		allFiles: allFiles,
		dirs:     dirs,
		allDirs:  allDirs,
		info:     info,
	}
}

func (d *Dir) String() string {
	return d.path
}

// Name 文件夹名称
func (d *Dir) Name() string {
	return d.info.Name()
}

func (d *Dir) Files() []*File {
	return d.files
}

func (d *Dir) Dirs() []*Dir {
	return d.dirs
}

func (d *Dir) AllFiles() []*File {
	return d.allFiles
}

func (d *Dir) AllDirs() []*Dir {
	return d.allDirs
}

func walkFiles(rootpath string) ([]*File, []*Dir) {
	var files []*File
	var dirs []*Dir
	var path string

	infos, err := ioutil.ReadDir(rootpath)
	if err != nil {
		panic(err)
	}

	for _, info := range infos {
		path = rootpath + "/" + info.Name()
		if !info.IsDir() {
			files = append(files, NewFile(path, info))
		} else {
			dirs = append(dirs, NewDir(path, info))
		}
	}
	return files, dirs
}

func walkAllFiles(rootpath string) ([]*File, []*Dir) {
	var files []*File
	var dirs []*Dir
	err := filepath.Walk(rootpath, visit(&files, &dirs, rootpath))
	if err != nil {
		panic(err)
	}

	return files, dirs
}

func visit(files *[]*File, dirs *[]*Dir, rootpath string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// walk will include rootpath
		if rootpath == path {
			return nil
		}

		if !info.IsDir() {
			*files = append(*files, NewFile(path, info))
		} else {
			*dirs = append(*dirs, NewDir(path, info))
		}
		return nil
	}
}
