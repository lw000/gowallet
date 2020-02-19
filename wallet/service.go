package wallet

import (
	log "github.com/sirupsen/logrus"
	"gowallet/dao/service"
	"sync"
)

type walletService struct {
	sync.RWMutex
	custWallets map[int]*CustWallet
}

var (
	__wallet_service      *walletService
	__wallet_service_once sync.Once
)

func WalletService() *walletService {
	__wallet_service_once.Do(func() {
		__wallet_service = &walletService{
			custWallets: make(map[int]*CustWallet),
		}
	})
	return __wallet_service
}

// 加载钱包
func (w *walletService) LoadWallets() error {
	serv := service.CustWalletDaoService{}
	wallets, err := serv.Query()
	if err != nil {
		log.Error(err)
		return err
	}

	for _, wallet := range wallets {
		custWallet, exists := w.custWallets[wallet.UserId]
		if !exists {
			custWallet = NewCustWallet(wallet.UserId)
			err = custWallet.LoadConfig()
			if err != nil {
				log.Error(err)
				return err
			}
			w.addCustWallet(custWallet)
		}
		custWallet.AddWallet(wallet)
	}
	return nil
}

// 存储钱包
func (w *walletService) StoreWallets() error {
	for _, custWallet := range w.custWallets {
		custWallet.StoreWallet()
	}
	return nil
}

func (w *walletService) addCustWallet(custWallet *CustWallet) {
	w.Lock()
	_, exists := w.custWallets[custWallet.Userid()]
	if !exists {
		w.custWallets[custWallet.Userid()] = custWallet
		w.Unlock()
		return
	}
	w.Unlock()
}
