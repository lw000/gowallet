package service

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"gowallet/dao/database"
	"gowallet/dao/mem"
	"gowallet/dao/table"
	"gowallet/errors"
)

type WalletDaoService struct {
}

func (serv *WalletDaoService) Preload() error {
	wallets, err := serv.Select()
	if err != nil {
		log.Error(err)
		return err
	}

	for _, wallet := range wallets {
		ca := mem.WalletCacheService{WalletId: wallet.WalletId}
		if err = ca.Save(wallet); err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

// 查询渠道信息
func (serv *WalletDaoService) Query(walletId int) (table.TWallet, error) {
	var (
		err    error
		wallet table.TWallet
	)

	// 查询缓存数据
	ca := mem.WalletCacheService{WalletId: walletId}
	wallet, err = ca.Load()
	if err != nil {
		log.WithFields(log.Fields{"walletId": walletId}).Error(err)
	}

	if wallet.WalletId == 0 {
		// 查询数据库
		wallet, err = serv.SelectWith(walletId)
		if err != nil {
			log.WithFields(log.Fields{"walletId": walletId}).Error(err)
			return table.TWallet{}, err
		}

		// 更新缓存
		if err = ca.Save(wallet); err != nil {
			log.WithFields(log.Fields{"walletId": walletId}).Error(err)
		}
	}

	return wallet, nil
}

// SelectWith
func (serv *WalletDaoService) SelectWith(walletId int) (table.TWallet, error) {
	query := `SELECT
				wallet_id,
				wallet_name,
				wallet_status
			FROM
				wallet
			WHERE
				wallet_id =?;`
	row := database.GetMysql(database.DB_WALLET).DB().QueryRow(query, walletId)
	var wallet table.TWallet
	err := row.Scan(
		&wallet.WalletId,
		&wallet.WalletName,
		&wallet.WalletStatus,
	)
	if err != nil && err != sql.ErrNoRows {
		log.WithFields(log.Fields{"wallet": walletId}).Error(err)
		return table.TWallet{}, err
	}

	if err == sql.ErrNoRows {
		return table.TWallet{}, errors.New(0, "未查询到钱包配置", "")
	}

	return wallet, nil
}

func (serv *WalletDaoService) Select() ([]table.TWallet, error) {
	query := `SELECT
				wallet_id,
				wallet_name,
				wallet_status
			FROM
				wallet;`
	rows, err := database.GetMysql(database.DB_WALLET).DB().Query(query)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	wallets := make([]table.TWallet, 0, 10)
	for rows.Next() {
		var wallet table.TWallet
		err = rows.Scan(
			&wallet.WalletId,
			&wallet.WalletName,
			&wallet.WalletStatus,
		)
		if err == nil {
			wallets = append(wallets, wallet)
		} else {
			log.Error(err)
		}
	}
	return wallets, nil
}
