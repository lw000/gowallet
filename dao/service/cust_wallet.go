package service

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"gowallet/dao/database"
	"gowallet/dao/table"
	"gowallet/errors"
)

type CustWalletDaoService struct {
}

// 查询钱包信息
func (serve *CustWalletDaoService) Query() ([]table.TCustWallet, error) {
	return serve.query()
}

// 查询钱包信息
func (serve *CustWalletDaoService) QueryBy(userId int) ([]table.TCustWallet, error) {
	return serve.queryBy(userId)
}

// 查询钱包信息
func (serve *CustWalletDaoService) QueryByWalletId(userId int, walletId int) (table.TCustWallet, error) {
	return serve.queryByWalletId(userId, walletId)
}

// queryByWalletId
func (serve *CustWalletDaoService) queryByWalletId(userId int, walletId int) (table.TCustWallet, error) {
	query := `SELECT
       			user_id,
				wallet_id,
				wallet_name,
       			amount,
       			frozen,
				status
			FROM
				cust_wallet
			WHERE
				user_id =? AND wallet_id =?;`
	row := database.GetMysql(database.DB_WALLET).DB().QueryRow(query, userId, walletId)
	var wallet table.TCustWallet
	err := row.Scan(
		&wallet.UserId,
		&wallet.WalletId,
		&wallet.WalletName,
		&wallet.Amount,
		&wallet.Frozen,
		&wallet.Status,
	)
	if err != nil && err != sql.ErrNoRows {
		log.WithFields(log.Fields{"userId": userId}).Error(err)
		return table.TCustWallet{}, err
	}

	if err == sql.ErrNoRows {
		return table.TCustWallet{}, errors.New(0, "未查询到用户钱包", "")
	}

	return wallet, nil
}

// queryBy
func (serve *CustWalletDaoService) queryBy(userId int) ([]table.TCustWallet, error) {
	query := `SELECT
       			user_id,
				wallet_id,
				wallet_name,
       			amount,
       			frozen,
				status
			FROM
				cust_wallet
			WHERE
				user_id =?;`
	rows, err := database.GetMysql(database.DB_WALLET).DB().Query(query)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	wallets := make([]table.TCustWallet, 0, 10)
	for rows.Next() {
		var wallet table.TCustWallet
		err = rows.Scan(
			&wallet.UserId,
			&wallet.WalletId,
			&wallet.WalletName,
			&wallet.Amount,
			&wallet.Frozen,
			&wallet.Status,
		)
		if err == nil {
			wallets = append(wallets, wallet)
		} else {
			log.Error(err)
		}
	}

	return wallets, nil
}

// query
func (serve *CustWalletDaoService) query() ([]table.TCustWallet, error) {
	query := `SELECT
				user_id,
				wallet_id,
				wallet_name,
       			amount,
       			frozen,
				status
			FROM
				cust_wallet;`
	rows, err := database.GetMysql(database.DB_WALLET).DB().Query(query)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	wallets := make([]table.TCustWallet, 0, 10)
	for rows.Next() {
		var wallet table.TCustWallet
		err = rows.Scan(
			&wallet.UserId,
			&wallet.WalletId,
			&wallet.WalletName,
			&wallet.Amount,
			&wallet.Frozen,
			&wallet.Status,
		)
		if err == nil {
			wallets = append(wallets, wallet)
		} else {
			log.Error(err)
		}
	}

	return wallets, nil
}
