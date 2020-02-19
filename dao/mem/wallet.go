package mem

import (
	"fmt"
	"gowallet/cache"
	"gowallet/dao/table"
)

// 钱包缓存对象 ...
type WalletCacheService struct {
	WalletId int
}

// key 缓存KEY
func (serve *WalletCacheService) key() string {
	key := fmt.Sprintf("wallet:%d", serve.WalletId)
	return key
}

// Exists ...
func (serve *WalletCacheService) Exists() bool {
	exist := cache.CommonCacheService().Exist(serve.key())
	return exist
}

// Load 缓存中读取数据
func (serve *WalletCacheService) Load() (table.TWallet, error) {
	v, err := cache.CommonCacheService().Get(serve.key())
	if err != nil {
		return table.TWallet{}, err
	}
	wallet := v.(table.TWallet)
	return wallet, nil
}

// Save 保存数据到Cache中
func (serve *WalletCacheService) Save(wallet table.TWallet) error {
	if err := cache.CommonCacheService().Set(serve.key(), wallet); err != nil {
		return err
	}
	return nil
}

// Clear ...
func (serve *WalletCacheService) Clear() bool {
	ok := cache.CommonCacheService().Remove(serve.key())
	return ok
}
