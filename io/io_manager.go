package io

const FileMod = 0644

type Manager interface {
	Read(b []byte, off int64) (int, error)
	Write(b []byte) (int, error)
	Sync() error
	Close() error
}
