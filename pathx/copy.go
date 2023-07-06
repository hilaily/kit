package pathx

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

// CopyFile copies a single file from src to dst
func CopyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer func() {
		_ = srcfd.Close()
	}()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer func() {
		_ = dstfd.Close()
	}()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

// CopyDir copies a whole directory recursively
func CopyDir(src string, dst string) error {
	var err error
	var fds []os.DirEntry
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = os.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcfp, dstfp); err != nil {
				log.Println(err)
			}
		} else {
			if err = CopyFile(srcfp, dstfp); err != nil {
				log.Println(err)
			}
		}
	}
	return nil
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
