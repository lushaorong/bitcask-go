package io

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestFileIO_Close(t *testing.T) {
	path := filepath.Join(os.TempDir(), "test.data")
	io, err := NewFileIOManager(path)
	defer destroyFile(path)
	assert.Nil(t, err)
	assert.NotNil(t, io)

	err = io.Close()
	assert.Nil(t, err)
}

func TestFileIO_Read(t *testing.T) {
	path := filepath.Join(os.TempDir(), "test.data")
	io, err := NewFileIOManager(path)
	defer destroyFile(path)
	assert.Nil(t, err)
	assert.NotNil(t, io)

	_, err = io.Write([]byte("key-a"))
	assert.Nil(t, err)

	_, err = io.Write([]byte("key-b"))
	assert.Nil(t, err)

	b1 := make([]byte, 5)
	n, err := io.Read(b1, 0)
	assert.Equal(t, 5, n)
	assert.Equal(t, []byte("key-a"), b1)

	b2 := make([]byte, 5)
	n, err = io.Read(b2, 5)
	assert.Equal(t, 5, n)
	assert.Equal(t, []byte("key-b"), b2)
}

func TestFileIO_Sync(t *testing.T) {
	path := filepath.Join(os.TempDir(), "test.data")
	defer destroyFile(path)

	io, err := NewFileIOManager(path)
	assert.Nil(t, err)
	assert.NotNil(t, io)

	err = io.Sync()
	assert.Nil(t, err)
}

func TestFileIO_Write(t *testing.T) {
	path := filepath.Join(os.TempDir(), "test.data")
	defer destroyFile(path)

	io, err := NewFileIOManager(path)
	assert.Nil(t, err)
	assert.NotNil(t, io)

	n, err := io.Write([]byte(""))
	assert.Equal(t, 0, n)
	assert.Nil(t, err)

	n, err = io.Write([]byte("hello world"))
	assert.Equal(t, 11, n)
	assert.Nil(t, err)

	n, err = io.Write([]byte("hello"))
	assert.Equal(t, 5, n)
	assert.Nil(t, err)
}

func destroyFile(name string) {
	if err := os.RemoveAll(name); err != nil {
		panic(err)
	}
}

func TestNewFileIOManager(t *testing.T) {
	path := filepath.Join(os.TempDir(), "test.data")
	defer destroyFile(path)

	io, err := NewFileIOManager(path)
	assert.Nil(t, err)
	assert.NotNil(t, io)
}
