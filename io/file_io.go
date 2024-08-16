package io

import "os"

type FileIO struct {
	fd *os.File
}

func NewFileIOManager(fileName string) (*FileIO, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, FileMod)
	if err != nil {
		return nil, err
	}
	return &FileIO{file}, nil
}

func (file *FileIO) Read(b []byte, index int64) (int, error) {
	return file.fd.ReadAt(b, index)
}

func (file *FileIO) Write(b []byte) (int, error) {
	return file.fd.Write(b)
}

func (file *FileIO) Sync() error {
	return file.fd.Sync()
}

func (file *FileIO) Close() error {
	return file.fd.Close()
}
