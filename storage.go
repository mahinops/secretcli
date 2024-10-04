package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	if fileName[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err == nil {
			fileName = filepath.Join(homeDir, fileName[2:])
		}
	}
	return &Storage[T]{FileName: fileName}
}

func (s *Storage[T]) Save(data T) error {
	dir := filepath.Dir(s.FileName)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, data)
}
