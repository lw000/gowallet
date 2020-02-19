package service

import (
	log "github.com/sirupsen/logrus"
	"gowallet/dao/database"
	"gowallet/dao/mem"
	"gowallet/dao/table"
)

type TradeInfoDaoService struct {
}

// 查询钱包信息
func (serv *TradeInfoDaoService) QueryWith(userId int) ([]table.TTradeInfo, error) {
	var (
		err        error
		tradeInfos []table.TTradeInfo
	)

	// 查询缓存数据
	ca := mem.TradeInfoCacheService{UserId: userId}
	tradeInfos, err = ca.Load()
	if err != nil {
		log.WithFields(log.Fields{"userId": userId}).Error(err)
	}

	if len(tradeInfos) == 0 {
		// 查询数据库
		tradeInfos, err = serv.Select(userId)
		if err != nil {
			log.WithFields(log.Fields{"userId": userId}).Error(err)
			return nil, err
		}

		// 更新缓存
		if err = ca.Save(tradeInfos...); err != nil {
			log.WithFields(log.Fields{"userId": userId}).Error(err)
		}
	}

	return tradeInfos, nil
}

// SelectWith
func (serv *TradeInfoDaoService) Select(userId int) ([]table.TTradeInfo, error) {
	query := `SELECT
				user_id,
				fwallet_id,
				twallet_id,
       			amount,
       			tradecode,
				tradetime,
       			tradestatus
			FROM
				trade_info
			WHERE user_id =?;`
	rows, err := database.GetMysql(database.DB_WALLET).DB().Query(query, userId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	tradeInfos := make([]table.TTradeInfo, 0, 10)
	for rows.Next() {
		var tradeInfo table.TTradeInfo
		err = rows.Scan(
			&tradeInfo.UserId,
			&tradeInfo.FwalletId,
			&tradeInfo.TwalletId,
			&tradeInfo.Amount,
			&tradeInfo.TradeCode,
			&tradeInfo.TradeTime,
			&tradeInfo.TradeStatus,
		)
		if err == nil {
			tradeInfos = append(tradeInfos, tradeInfo)
		} else {
			log.Error(err)
		}
	}
	return tradeInfos, nil
}
