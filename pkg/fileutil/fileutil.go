package fileutil

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func SaveFile(
	r io.Reader,
	dst string,
) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, r)
	return err
}

func SaveMultipartFile(
	file *multipart.FileHeader,
	dst string,
) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	return SaveFile(src, dst)
}

type MultipartFileInfo struct {
	FileName string
	MimeType string
	Size     float64
}

func NewMultipartFileInfo(file *multipart.FileHeader) (*MultipartFileInfo, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Read first 512 bytes to detect MIME type
	buffer := make([]byte, 512)
	n, err := f.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, err
	}

	mimeType := http.DetectContentType(buffer[:n])

	info := &MultipartFileInfo{
		FileName: file.Filename,
		MimeType: mimeType,
		Size:     float64(file.Size),
	}

	return info, nil
}
