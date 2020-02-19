package service

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"gowallet/dao/database"
	"gowallet/dao/table"
	"gowallet/errors"
)

// 查询用户钱包配置
type CustWalletPwdDaoService struct {
}

// 查询用户钱包配置
func (serv *CustWalletPwdDaoService) Query() ([]table.TCustWalletPwd, error) {
	return serv.query()
}

// 查询用户钱包配置
func (serv *CustWalletPwdDaoService) QueryBy(userId int) (table.TCustWalletPwd, error) {
	return serv.queryBy(userId)
}

// 查询用户钱包配置
func (serv *CustWalletPwdDaoService) queryBy(userId int) (table.TCustWalletPwd, error) {
	query := `SELECT
       			user_id,
				pwd,
				pwdtype,
       			lockstatus,
       			lockdate,
				lockexpiredate
			FROM
				cust_wallet_pwd
			WHERE
				user_id =?;`
	row := database.GetMysql(database.DB_WALLET).DB().QueryRow(query, userId)
	var custWalletPwd table.TCustWalletPwd
	err := row.Scan(
		&custWalletPwd.UserId,
		&custWalletPwd.Pwd,
		&custWalletPwd.Pwdtype,
		&custWalletPwd.Lockstatus,
		&custWalletPwd.Lockdate,
		&custWalletPwd.LockExpiredate,
	)
	if err != nil && err != sql.ErrNoRows {
		log.WithFields(log.Fields{"userId": userId}).Error(err)
		return table.TCustWalletPwd{}, err
	}

	if err == sql.ErrNoRows {
		return table.TCustWalletPwd{}, errors.New(0, "未查询到用户钱包", "")
	}

	return custWalletPwd, nil
}

// 查询用户钱包配置
func (serv *CustWalletPwdDaoService) query() ([]table.TCustWalletPwd, error) {
	query := `SELECT
				user_id,
				pwd,
				pwdtype,
       			lockstatus,
       			lockdate,
				lockexpiredate
			FROM
				cust_wallet_pwd;`
	rows, err := database.GetMysql(database.DB_WALLET).DB().Query(query)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	custWalletPwds := make([]table.TCustWalletPwd, 0, 10)
	for rows.Next() {
		var custWalletPwd table.TCustWalletPwd
		err = rows.Scan(
			&custWalletPwd.UserId,
			&custWalletPwd.Pwd,
			&custWalletPwd.Pwdtype,
			&custWalletPwd.Lockstatus,
			&custWalletPwd.Lockdate,
			&custWalletPwd.LockExpiredate,
		)
		if err == nil {
			custWalletPwds = append(custWalletPwds, custWalletPwd)
		} else {
			log.Error(err)
		}
	}
	return custWalletPwds, nil
}
