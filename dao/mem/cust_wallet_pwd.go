package mem

import (
	"fmt"
	"gowallet/cache"
	"gowallet/dao/table"
)

// 用户钱包配置缓存对象 ...
type CustWalletPwdCacheService struct {
	UserId int
}

// key 缓存KEY
func (serve *CustWalletPwdCacheService) key() string {
	key := fmt.Sprintf("cust_wallet_pwd:%d", serve.UserId)
	return key
}

// Exists ...
func (serve *CustWalletPwdCacheService) Exists() bool {
	exist := cache.CommonCacheService().Exist(serve.key())
	return exist
}

// Load 缓存中读取数据
func (serve *CustWalletPwdCacheService) Load() (table.TCustWalletPwd, error) {
	v, err := cache.CommonCacheService().Get(serve.key())
	if err != nil {
		return table.TCustWalletPwd{}, err
	}
	wallet := v.(table.TCustWalletPwd)
	return wallet, nil
}

// Save 保存数据到Cache中
func (serve *CustWalletPwdCacheService) Save(wallet table.TCustWalletPwd) error {
	if err := cache.CommonCacheService().Set(serve.key(), wallet); err != nil {
		return err
	}
	return nil
}

// Clear ...
func (serve *CustWalletPwdCacheService) Clear() bool {
	ok := cache.CommonCacheService().Remove(serve.key())
	return ok
}
