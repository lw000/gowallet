package service

import (
	log "github.com/sirupsen/logrus"
)

func Test() {
	{
		serv := WalletDaoService{}
		err := serv.Preload()
		if err != nil {
			log.Error(err)
			return
		}
	}

	{
		serv := CustWalletDaoService{}
		data, err := serv.Query()
		if err != nil {
			log.Error(err)
			return
		}
		if len(data) > 0 {

		}
	}

	{
		serv := CustWalletPwdDaoService{}
		data, err := serv.Query()
		if err != nil {
			log.Error(err)
			return
		}
		if len(data) > 0 {

		}
	}
}
