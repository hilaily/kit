package pathx

import (
	"fmt"
	"os"
	"path/filepath"
)

// IsExist if a file or path is exist
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func ChmodR(name string, mode os.FileMode) error {
	return filepath.Walk(name, func(path string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chmod(path, mode)
		}
		return err
	})
}

func ChownR(path string, uid, gid int) error {
	return filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chown(name, uid, gid)
		}
		return err
	})
}

// AppendFile writes data to a file named by filename.
// If the file does not exist, WriteFile creates it with permissions perm
// (before umask).
func AppendFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, perm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// NewEntity ...
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
func (n *Entity) IsDir() bool {
	n.initStat()
	return n.isDir
}

// IsExist ...
func (n *Entity) Exist() bool {
	n.initStat()
	return n.isExist
}

// Reload represent reread file from os
func (n *Entity) Reload() {
	n.isInit = false
	n.initStat()
}

func (n *Entity) initStat() {
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

// CleanDir ...
func CleanDir(dir string) error {
	err := os.RemoveAll(dir)
	if err != nil {
		return fmt.Errorf("remove %s %w", dir, err)
	}
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return fmt.Errorf("mkdir %s %w", dir, err)
	}
	return nil
}
