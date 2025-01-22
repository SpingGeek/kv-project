// Package data @Author Gopher
// @Date 2025/1/22 15:26:00
// @Desc 内存数据记录
package data

// LogRecordPos 数据内存索引，用于快速定位数据位置，主要是描述数据在磁盘上的位置
type LogRecordPos struct {
	Fid    uint32 // 文件 id，表示将数据存储到了哪个文件当中
	Offset int64  // 偏移量，表示将数据存储到数据文件中哪个位置
}
