package finder

import (
	"os"
)

// NewEntiry ...
func NewEntity(path string) *Entity {
	n := &Entity{path: path}
	return n
}

// Entity represent a file or folder
type Entity struct {
	path     string
	fileInfo *os.FileInfo
	isExist  bool
	isDir    bool
	isInit   bool
}

// IsDir ...
func (n Entity) IsDir() bool {
	n.initStat()
	return n.isDir
}

// IsExist ...
func (n Entity) Exist() bool {
	n.initStat()
	return n.isExist
}

// Reload represent reread file from os
func (n Entity) Reload() {
	n.isInit = false
	n.initStat()
}

func (n Entity) initStat() {
	if !n.isInit {
		stat, err := os.Stat(n.path)
		if err != nil {
			if os.IsExist(err) {
				n.isExist = true
			}
			n.isExist = false
		} else {
			n.isDir = stat.IsDir()
			n.isExist = true
			n.fileInfo = &stat
		}
		n.isInit = true
	}
}
