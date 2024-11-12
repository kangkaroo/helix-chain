package storage

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

// LevelDB 管理结构体
type LevelDB struct {
	db *leveldb.DB
}

// 打开或创建 LevelDB 数据库
func NewLevelDB(dbPath string) (*LevelDB, error) {
	// 打开 LevelDB 数据库
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return nil, fmt.Errorf("无法打开数据库: %v", err)
	}
	return &LevelDB{db: db}, nil
}

// 关闭数据库
func (ldb *LevelDB) Close() error {
	return ldb.db.Close()
}

// 设置值到 LevelDB
func (ldb *LevelDB) Put(key []byte, value []byte) error {
	return ldb.db.Put(key, value, nil)
}

// 从 LevelDB 获取值
func (ldb *LevelDB) Get(key []byte) ([]byte, error) {
	return ldb.db.Get(key, nil)
}

// 检查键是否存在
func (ldb *LevelDB) Has(key []byte) (bool, error) {
	return ldb.db.Has(key, nil)
}

// 删除值
func (ldb *LevelDB) Delete(key []byte) error {
	return ldb.db.Delete(key, nil)
}

// 获取所有键值对（调试使用）
func (ldb *LevelDB) GetAll() (map[string][]byte, error) {
	iter := ldb.db.NewIterator(nil, nil)
	defer iter.Release()

	result := make(map[string][]byte)
	for iter.Next() {
		result[string(iter.Key())] = iter.Value()
	}

	if err := iter.Error(); err != nil {
		return nil, err
	}

	return result, nil
}
