package utils

import (
	"crypto/sha1"
	"fmt"
	gouuid "github.com/satori/go.uuid"
	"io"
)

func GetFileSha1(f io.Reader) (string, error) {
	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("% x", h.Sum(nil)), nil
}

func GenerateUUID() string {
	return fmt.Sprintf("%s", gouuid.NewV4())
}
