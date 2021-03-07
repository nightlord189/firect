package hooks

import (
	"fmt"
	"os"
)

//FileHook - hook for printing logs to file
type FileHook struct {
	File *os.File
}

//NewFileHook - constructor
//path - path to log file
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

//LogStr - log just string
func (h *FileHook) LogStr(str string) error {
	_, err := h.File.WriteString(str + "\n")
	return err
}

//Log - log map
func (h *FileHook) Log(fields map[string]interface{}) error {
	_, err := h.File.WriteString(fmt.Sprintf("%v\n", fields))
	return err
}
