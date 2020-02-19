package wallet

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gowallet/dao/service"
	"gowallet/dao/table"
	"sync"
	"time"
)

type CustWallet struct {
	sync.RWMutex
	userid         int       // 用户ID
	pwd            string    // 钱包密码
	lockStatus     int       // 锁定状态
	lockDate       time.Time // 锁定日期
	lockExpireDate time.Time // 锁定过期日期
	wallets        map[int]*Wallet
}

func NewCustWallet(userid int) *CustWallet {
	return &CustWallet{
		userid:  userid,
		wallets: make(map[int]*Wallet),
	}
}

func (w *CustWallet) LoadConfig() error {
	serv := service.CustWalletPwdDaoService{}
	conf, err := serv.QueryBy(w.userid)
	if err != nil {
		log.Error(err)
		return err
	}

	w.Lock()
	// 钱包密码
	w.pwd = conf.Pwd

	// 钱包状态
	w.lockStatus = conf.Lockstatus

	// 钱包锁定日期
	w.lockDate = conf.Lockdate

	// 钱包锁定过期日期
	w.lockExpireDate = conf.LockExpiredate

	ts := w.lockExpireDate.Sub(w.lockDate)
	if ts > 0 {
		log.WithField("锁定时间", fmt.Sprintf("%v", ts)).Info("钱包锁定")
	} else {

	}

	w.Unlock()

	return nil
}

// 添加钱包
func (w *CustWallet) AddWallet(wallet table.TCustWallet) {
	w.Lock()
	_, exists := w.wallets[wallet.WalletId]
	if !exists {
		w.wallets[wallet.WalletId] = &Wallet{
			walletId:   wallet.WalletId,
			walletName: wallet.WalletName,
			amount:     wallet.Amount,
			frozen:     wallet.Frozen,
			status:     wallet.Status,
		}
		w.Unlock()
		return
	}
	w.Unlock()
}

// 存储钱包
func (w *CustWallet) StoreWallet() {
	w.Lock()
	defer w.Unlock()

}

func (w *CustWallet) Userid() int {
	return w.userid
}

func (w *CustWallet) Wallets() []*Wallet {
	w.Lock()
	var wallets []*Wallet
	for _, wallet := range w.wallets {
		wallets = append(wallets, wallet)
	}
	w.Unlock()
	return wallets
}

func (w *CustWallet) Wallet(walletId int) *Wallet {
	w.Lock()
	wallet, exists := w.wallets[walletId]
	if exists {
		w.Unlock()
		return wallet
	}
	w.Unlock()
	return nil
}
