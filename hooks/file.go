package hooks

import (
	"fmt"
	"os"
)

type FileHook struct {
	File *os.File
}

func NewFileHook(path string) (*FileHook, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	hook := FileHook{
		File: f,
	}
	return &hook, nil
}

func (h *FileHook) LogStr(str string) error {
	_, err := h.File.WriteString(str)
	return err
}

func (h *FileHook) Log(fields map[string]interface{}) error {
	_, err := h.File.WriteString(fmt.Sprintf("%v", fields))
	return err
}
