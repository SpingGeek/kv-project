// Package index @Author Gopher
// @Date 2025/1/22 15:33:00
// @Desc BTree 索引模块
package index

import (
	"bitcask-go/data"
	"sync"

	"github.com/google/btree"
)

// BTree 索引，主要是封装了 Google 的 btree 库
// https://github.com/google/btree
type BTree struct {
	tree *btree.BTree  // 引入 BTree 索引
	lock *sync.RWMutex // 由于在多个 goroutines 写操作是不安全的，则实现加锁保护
}

// NewBTree 初始化 BTree 索引结构
func NewBTree() *BTree {
	return &BTree{
		// 提供一个初始化参数，让用户来进行选择
		tree: btree.New(32),
		lock: new(sync.RWMutex),
	}
}

// Put 向索引中存储 key 对应的数据位置信息，写入数据
func (bt *BTree) Put(key []byte, pos *data.LogRecordPos) bool {
	it := &Item{key: key, pos: pos}
	bt.lock.Lock()
	bt.tree.ReplaceOrInsert(it)
	bt.lock.Unlock()
	return true
}

// Get 根据 key 获取对应的索引数据位置信息
func (bt *BTree) Get(key []byte) *data.LogRecordPos {
	it := &Item{key: key}
	btreeItem := bt.tree.Get(it)
	if btreeItem == nil {
		return nil
	}
	return btreeItem.(*Item).pos
}

// Delete 根据 key 删除对应的索引数据位置信息
func (bt *BTree) Delete(key []byte) bool {
	it := &Item{key: key}
	bt.lock.Lock()
	oldItem := bt.tree.Delete(it)
	// 这个地方之前在加锁的过程中是忘记释放的，所以在 test_delete 出现了 Bug
	bt.lock.Unlock()
	if oldItem == nil {
		return false
	}
	return true
}
