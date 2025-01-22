// Package fio @Author Gopher
// @Date 2025/1/22 17:13:00
// @Desc IO 管理器接口定义
package fio

// DataFilePerm 文件权限定义常量
const DataFilePerm = 0644

// IOManger 抽象 IO 管理接口，可以接入不同的 IO 类型，目前支持标准文件 IO
type IOManger interface {

	// Read 从文件的给定位置读取对应的数据
	Read([]byte, int64) (int, error)

	// Write 写入字节数组到文件中
	Write([]byte) (int, error)

	// Sync 将内存中的数据同步到磁盘中，持久化数据
	Sync() error

	// Close 关闭文件
	Close() error
}
