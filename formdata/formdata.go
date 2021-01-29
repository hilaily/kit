package formdata

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
)

// NewFormBody represent create a http form data request
// @param params represent fields
// @param files represent update files
func NewFormBody(params map[string]string, files []*FileInfo) (contentType string, body io.Reader, err error) {
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)

	defer writer.Close()

	for _, v := range files {
		fileWriter, err := writer.CreateFormFile(v.Fieldname, v.Filename)
		if err != nil {
			err = fmt.Errorf("write file field, key: %s, val: %s, %w", v.Fieldname, v.Filename, err)
			return
		}
		_, err = io.Copy(fileWriter, v.Data)
		if err != nil {
			err = fmt.Errorf("write file, key: %s, val: %s, %w", v.Fieldname, v.Filename, err)
			return
		}
	}

	for k, v := range params {
		err = writer.WriteField(k, v)
		if err != nil {
			err = fmt.Errorf("write field, key: %s, val: %s, %w", k, v, err)
			return
		}
	}

	return writer.FormDataContentType(), buf, nil
}

// FileInfo for update file info
type FileInfo struct {
	Fieldname string
	Filename  string
	Data      io.Reader
}
