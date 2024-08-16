package index

import (
	"bitcask-go/data"
)

type Indexer interface {
	Put(key []byte, value *data.LogRecordPos) bool
	Get(key []byte) *data.LogRecordPos
	Delete(key []byte) bool
}
