package common

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/google/uuid"
)

func SaveFile(file *multipart.FileHeader, dir string) (string, error) {
	randFileName := uuid.New()
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}
	fileNameArr := strings.Split(file.Filename, ".")
	ext := fileNameArr[len(fileNameArr)-1]
	fileName := fmt.Sprintf("%s.%s", randFileName, ext)
	dst := fmt.Sprintf("%s/%s", dir, fileName)

	src, err := file.Open()
	if err != nil {
		return "", nil
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
