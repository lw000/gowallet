package mem

import (
	"fmt"
	"gowallet/cache"
	"gowallet/dao/table"
)

// 交易信息缓存对象 ...
type TradeInfoCacheService struct {
	UserId int
}

// key 缓存KEY
func (serve *TradeInfoCacheService) key() string {
	key := fmt.Sprintf("trade_info:%d", serve.UserId)
	return key
}

// Exists ...
func (serve *TradeInfoCacheService) Exists() bool {
	exist := cache.CommonCacheService().Exist(serve.key())
	return exist
}

// Load 缓存中读取数据
func (serve *TradeInfoCacheService) Load() ([]table.TTradeInfo, error) {
	v, err := cache.CommonCacheService().Get(serve.key())
	if err != nil {
		return nil, err
	}
	tradeInfo := v.([]table.TTradeInfo)
	return tradeInfo, nil
}

// Save 保存数据到Cache中
func (serve *TradeInfoCacheService) Save(tradeInfo ...table.TTradeInfo) error {
	if err := cache.CommonCacheService().Set(serve.key(), tradeInfo); err != nil {
		return err
	}
	return nil
}

// Clear ...
func (serve *TradeInfoCacheService) Clear() bool {
	ok := cache.CommonCacheService().Remove(serve.key())
	return ok
}
