package dao

import (
	"fmt"
	"gowallet/dao/service"
	"gowallet/dao/table"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/singleflight"
)

type ProloadDaoServiceInterface interface {
	Preload() error
}

var (
	singleflightWallet singleflight.Group
)

// 查询游戏地址
func QueryWallet(walletId int) ([]table.TWallet, error) {
	v, err, sh := singleflightWallet.Do(fmt.Sprintf("wallet_%d", walletId), func() (i interface{}, e error) {
		serv := service.WalletDaoService{}
		return serv.Query(walletId)
	})

	if err != nil {
		log.Error(err)
		return nil, err
	}

	if sh {
	}

	return v.([]table.TWallet), nil
}
