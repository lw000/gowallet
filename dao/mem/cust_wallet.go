package mem

import (
	"fmt"
	"gowallet/cache"
	"gowallet/dao/table"
)

// 用户钱包缓存对象 ...
type CustWalletCacheService struct {
	UserId   int
	WalletId int
}

// key 缓存KEY
func (serve *CustWalletCacheService) key() string {
	key := fmt.Sprintf("cust_wallet:%d:%d", serve.UserId, serve.WalletId)
	return key
}

// Exists ...
func (serve *CustWalletCacheService) Exists() bool {
	exist := cache.CommonCacheService().Exist(serve.key())
	return exist
}

// Load 缓存中读取数据
func (serve *CustWalletCacheService) Load() (table.TCustWallet, error) {
	v, err := cache.CommonCacheService().Get(serve.key())
	if err != nil {
		return table.TCustWallet{}, err
	}
	wallet := v.(table.TCustWallet)
	return wallet, nil
}

// Save 保存数据到Cache中
func (serve *CustWalletCacheService) Save(wallet table.TCustWallet) error {
	if err := cache.CommonCacheService().Set(serve.key(), wallet); err != nil {
		return err
	}
	return nil
}

// Clear ...
func (serve *CustWalletCacheService) Clear() bool {
	ok := cache.CommonCacheService().Remove(serve.key())
	return ok
}
