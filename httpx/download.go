package httpx

import (
	"io"
	"net/http"
)

func DownloadHandler(resp http.ResponseWriter, reader io.Reader, filename string) error {
	return DownloadHander(resp, reader, filename)
}

func DownloadHander(resp http.ResponseWriter, reader io.Reader, filename string) error {
	resp.Header().Add("Content-Description", "File Transfer")
	resp.Header().Add("Content-Disposition", "attachment; filename="+filename)
	resp.Header().Add("Content-Type", "application/octet-stream")
	_, err := io.Copy(resp, reader)
	return err
}
