package index

import (
	"bitcask-go/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBTree_Delete(t *testing.T) {
	indexer := NewBTree(32)

	a1 := indexer.Delete([]byte("a"))
	assert.False(t, a1)

	indexer.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 10})
	a2 := indexer.Delete([]byte("a"))
	assert.True(t, a2)
}

func TestBTree_Get(t *testing.T) {
	indexer := NewBTree(32)

	a1 := indexer.Get([]byte("a"))
	assert.True(t, a1 == nil)

	pos := &data.LogRecordPos{Fid: 1, Offset: 10}
	indexer.Put([]byte("a"), pos)
	a2 := indexer.Get([]byte("a"))
	assert.Equal(t, pos, a2)
}

func TestBTree_Put(t *testing.T) {
	indexer := NewBTree(32)

	a1 := indexer.Put([]byte("a"), &data.LogRecordPos{
		Fid:    1,
		Offset: 100,
	})
	assert.True(t, a1)

	a2 := indexer.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, a2)
}
