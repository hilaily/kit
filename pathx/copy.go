package pathx

import (
	"fmt"
	"io"
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
		return fmt.Errorf("stat %s %w", src, err)
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return fmt.Errorf("mkdir %s %w", dst, err)
	}

	if fds, err = os.ReadDir(src); err != nil {
		return fmt.Errorf("read dir %s %w", src, err)
	}
	for _, fd := range fds {
		err = copyOne(src, dst, fd)
		if err != nil {
			return err
		}
	}
	return nil
}

func copyOne(src, dst string, fd os.DirEntry) error {
	srcfp := path.Join(src, fd.Name())
	dstfp := path.Join(dst, fd.Name())

	var err error
	if fd.IsDir() {
		return CopyDir(srcfp, dstfp)
	}
	err = CopyFile(srcfp, dstfp)
	if err != nil {
		return fmt.Errorf("copy file %s - %s %w", srcfp, dstfp, err)
	}
	return nil
}
