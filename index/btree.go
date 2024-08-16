package index

import (
	"bitcask-go/data"
	"bytes"
	"github.com/google/btree"
	_ "github.com/google/btree"
	"sync"
)

type BTree struct {
	tree *btree.BTree
	mu   *sync.RWMutex
}

type BTreeItem struct {
	key   []byte
	value *data.LogRecordPos
}

func (t *BTreeItem) Less(than btree.Item) bool {
	return bytes.Compare(t.key, than.(*BTreeItem).key) == -1
}

func NewBTree(degree int) *BTree {
	if degree == 0 {
		degree = 32
	}
	return &BTree{
		tree: btree.New(degree),
		mu:   new(sync.RWMutex),
	}
}

func (bt *BTree) Put(key []byte, value *data.LogRecordPos) bool {
	bt.mu.Lock()
	defer bt.mu.Unlock()
	bt.tree.ReplaceOrInsert(&BTreeItem{
		key: key, value: value,
	})
	return true
}

func (bt *BTree) Get(key []byte) *data.LogRecordPos {
	item := bt.tree.Get(&BTreeItem{
		key: key,
	})
	if item == nil {
		return nil
	}
	return item.(*BTreeItem).value
}

func (bt *BTree) Delete(key []byte) bool {
	bt.mu.Lock()
	defer bt.mu.Unlock()
	if value := bt.tree.Delete(&BTreeItem{key: key}); value == nil {
		return false
	}
	return true
}
