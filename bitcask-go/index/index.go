// Package index @Author Gopher
// @Date 2025/1/22 15:23:00
// @Desc 通过相关数据结构对象来实现索引
package index

import (
	"bitcask-go/data"
	"bytes"

	"github.com/google/btree"
)

// Indexer 抽象索引接口，后续如果想要接入其他的数据结构，则直接实现这个接口即可
type Indexer interface {
	// Put 向索引中存储 key 对应的数据位置信息，写入数据
	Put(key []byte, pos *data.LogRecordPos) bool

	// Get 根据 key 获取对应的索引数据位置信息
	Get(key []byte) *data.LogRecordPos

	// Delete 根据 key 删除对应的索引数据位置信息
	Delete(key []byte) bool
}

// Item 索引项
type Item struct {
	key []byte
	pos *data.LogRecordPos
}

// Less 比较函数
func (ai *Item) Less(bi btree.Item) bool {
	return bytes.Compare(ai.key, bi.(*Item).key) == -1
}
